/**
 * BIN List Download
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

export async function get_bin_list_download(include-iso3, include-8digit) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (include-iso3) params.append("include-iso3", include-iso3);
      if (include-8digit) params.append("include-8digit", include-8digit);
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

export function createGetBinListDownloadTool() {
  return {
    definition: {
      name: 'get-bin-list-download',
      description: 'BIN List Download',
      inputSchema: {
        type: 'object',
        properties: {
          include-iso3: {
            type: 'boolean',
            description: 'Include ISO 3-letter country codes and ISO 3-letter currency codes in the data. These will be added to columns 10 and 11 respectively'
          },
          include-8digit: {
            type: 'boolean',
            description: 'Include 8-digit and higher BIN codes. This option includes all 6-digit BINs and all 8-digit and higher BINs (including some 9, 10 and 11 digit BINs where available)'
          }
        },
        required: []
      }
    },
    handler: get_bin_list_download
  };
}