/**
 * Convert
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

export async function get_convert(from-value, from-type, to-type) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (from-value) params.append("from-value", from-value);
      if (from-type) params.append("from-type", from-type);
      if (to-type) params.append("to-type", to-type);
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

export function createGetConvertTool() {
  return {
    definition: {
      name: 'get-convert',
      description: 'Convert',
      inputSchema: {
        type: 'object',
        properties: {
          from-value: {
            type: 'string',
            description: 'The value to convert from (e.g. 10.95)'
          },
          from-type: {
            type: 'string',
            description: 'The type of the value to convert from (e.g. USD)'
          },
          to-type: {
            type: 'string',
            description: 'The type to convert to (e.g. EUR)'
          }
        },
        required: ["from-value", "from-type", "to-type"]
      }
    },
    handler: get_convert
  };
}