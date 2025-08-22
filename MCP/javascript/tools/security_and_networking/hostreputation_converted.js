/**
 * Host Reputation
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

export async function get_host_reputation(host, zones, list-rating) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (host) params.append("host", host);
      if (zones) params.append("zones", zones);
      if (list-rating) params.append("list-rating", list-rating);
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

export function createGetHostReputationTool() {
  return {
    definition: {
      name: 'get-host-reputation',
      description: 'Host Reputation',
      inputSchema: {
        type: 'object',
        properties: {
          host: {
            type: 'string',
            description: 'An IP address, domain name, FQDN or URL. <br>If you supply a domain/URL it will be checked against the URI DNSBL lists'
          },
          zones: {
            type: 'string',
            description: 'Only check these DNSBL zones/hosts. Multiple zones can be supplied as comma-separated values'
          },
          list-rating: {
            type: 'number',
            description: 'Only check lists with this rating or better'
          }
        },
        required: ["host"]
      }
    },
    handler: get_host_reputation
  };
}