/**
 * IP Info
 */

import fs from 'fs';
import path from 'path';
import os from 'os';

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

export async function get_ip_info(ip, reverse-lookup) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (ip) params.append("ip", ip);
      if (reverse-lookup) params.append("reverse-lookup", reverse-lookup);
    const queryString = params.toString();
    const finalUrl = queryString ? `${url}?${queryString}` : url;
    
    const url = `${config.baseURL}/api/unknown`;
    
    const response = await fetch(finalUrl, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${config.bearerToken}`,
        'Accept': 'application/json'
      }
    });
    
    if (!response.ok) {
      return `Failed to format JSON: ${response.status} ${response.statusText}`;
    }
    
    try {
      const result = await response.json();
      return JSON.stringify(result, null, 2);
    } catch (e) {
      return await response.text();
    }
    
  } catch (error) {
    return `Request failed: ${error.message}`;
  }
}

export function createGetIpInfoTool() {
  return {
    definition: {
      name: 'get-ip-info',
      description: 'IP Info',
      inputSchema: {
        type: 'object',
        properties: {
          ip: {
            type: 'string',
            description: 'IPv4 or IPv6 address'
          },
          reverse-lookup: {
            type: 'boolean',
            description: 'Do a reverse DNS (PTR) lookup. This option can add extra delay to the request so only use it if you need it'
          }
        },
        required: ["ip"]
      }
    },
    handler: get_ip_info
  };
}