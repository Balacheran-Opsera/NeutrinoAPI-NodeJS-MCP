/**
 * MCP Server - JavaScript Implementation
 */

import { Server } from '@modelcontextprotocol/sdk/server/index.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import { ListToolsRequestSchema, CallToolRequestSchema } from '@modelcontextprotocol/sdk/types.js';
import fs from 'fs';
import path from 'path';
import os from 'os';

import { get_ua_lookup, createGetUaLookupTool } from './tools/data_tools/ualookup_converted.js';
import { get_email_validate, createGetEmailValidateTool } from './tools/data_tools/emailvalidate_converted.js';
import { get_phone_validate, createGetPhoneValidateTool } from './tools/data_tools/phonevalidate_converted.js';
import { get_bin_list_download, createGetBinListDownloadTool } from './tools/e_commerce/binlistdownload_converted.js';
import { get_bin_lookup, createGetBinLookupTool } from './tools/e_commerce/binlookup_converted.js';
import { get_convert, createGetConvertTool } from './tools/e_commerce/convert_converted.js';
import { get_ip_blocklist_download, createGetIpBlocklistDownloadTool } from './tools/security_and_networking/ipblocklistdownload_converted.js';
import { get_ip_probe, createGetIpProbeTool } from './tools/security_and_networking/ipprobe_converted.js';
import { get_domain_lookup, createGetDomainLookupTool } from './tools/security_and_networking/domainlookup_converted.js';
import { get_host_reputation, createGetHostReputationTool } from './tools/security_and_networking/hostreputation_converted.js';
import { get_ip_blocklist, createGetIpBlocklistTool } from './tools/security_and_networking/ipblocklist_converted.js';
import { get_email_verify, createGetEmailVerifyTool } from './tools/security_and_networking/emailverify_converted.js';
import { get_ip_info, createGetIpInfoTool } from './tools/geolocation/ipinfo_converted.js';
import { get_geocode_address, createGetGeocodeAddressTool } from './tools/geolocation/geocodeaddress_converted.js';
import { get_geocode_reverse, createGetGeocodeReverseTool } from './tools/geolocation/geocodereverse_converted.js';
import { get_hlr_lookup, createGetHlrLookupTool } from './tools/telephony/hlrlookup_converted.js';
import { get_verify_security_code, createGetVerifySecurityCodeTool } from './tools/telephony/verifysecuritycode_converted.js';
import { get_url_info, createGetUrlInfoTool } from './tools/www/urlinfo_converted.js';

// Create MCP server
const server = new Server({
  name: 'MCP Server',
  version: '1.0.0'
}, {
  capabilities: {
    tools: {}
  }
});

function getConfig() {
  const baseURL = process.env.API_BASE_URL;
  const bearerToken = process.env.API_BEARER_TOKEN;
  
  if (!baseURL || !bearerToken) {
    const configPath = path.join(os.homedir(), '.api', 'config.json');
    try {
      const configData = JSON.parse(fs.readFileSync(configPath, 'utf8'));
      return {
        baseURL: baseURL || configData.baseURL,
        bearerToken: bearerToken || configData.bearerToken
      };
    } catch (e) {
      throw new Error('Configuration not found. Please set API_BASE_URL and API_BEARER_TOKEN environment variables or create config file at ~/.api/config.json');
    }
  }
  
  return { baseURL, bearerToken };
}

// Register all tools
const tools = [
  createGetUaLookupTool(),
  createGetEmailValidateTool(),
  createGetPhoneValidateTool(),
  createGetBinListDownloadTool(),
  createGetBinLookupTool(),
  createGetConvertTool(),
  createGetIpBlocklistDownloadTool(),
  createGetIpProbeTool(),
  createGetDomainLookupTool(),
  createGetHostReputationTool(),
  createGetIpBlocklistTool(),
  createGetEmailVerifyTool(),
  createGetIpInfoTool(),
  createGetGeocodeAddressTool(),
  createGetGeocodeReverseTool(),
  createGetHlrLookupTool(),
  createGetVerifySecurityCodeTool(),
  createGetUrlInfoTool()
];

// List all available tools
function listToolsHandler() {
  return { tools: tools.map(tool => tool.definition) };
}

// Handle tool calls
function createCallToolHandler(toolMap) {
  return async (request) => {
    const { name, arguments: args } = request.params;
    
    const tool = toolMap.find(t => t.definition.name === name);
    if (!tool) {
      throw new Error(`Unknown tool: ${name}`);
    }
    
    try {
      const result = await tool.handler(args);
      return {
        content: [{
          type: 'text',
          text: result
        }]
      };
    } catch (error) {
      throw new Error(`Tool execution failed: ${error.message}`);
    }
  };
}

// Setup request handlers
server.setRequestHandler(ListToolsRequestSchema, listToolsHandler);
server.setRequestHandler(CallToolRequestSchema, createCallToolHandler(tools));

async function main() {
  try {
    const config = getConfig();
    console.error('MCP Server started successfully');
    
    const transport = new StdioServerTransport();
    await server.connect(transport);
  } catch (error) {
    console.error('Failed to start server:', error);
    process.exit(1);
  }
}

if (import.meta.url === `file://${process.argv[1]}`) {
  main().catch(console.error);
}