/**
 * BIN Lookup
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

export async function get_bin_lookup(bin-number, customer-ip) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (bin-number) params.append("bin-number", bin-number);
      if (customer-ip) params.append("customer-ip", customer-ip);
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

export function createGetBinLookupTool() {
  return {
    definition: {
      name: 'get-bin-lookup',
      description: 'BIN Lookup',
      inputSchema: {
        type: 'object',
        properties: {
          bin-number: {
            type: 'string',
            description: 'The BIN or IIN number. This is the first 6, 8 or 10 digits of a card number, use 8 (or more) digits for the highest level of accuracy'
          },
          customer-ip: {
            type: 'string',
            description: 'Pass in the customers IP address and we will return some extra information about them'
          }
        },
        required: ["bin-number"]
      }
    },
    handler: get_bin_lookup
  };
}