/**
 * Domain Lookup
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

export async function get_domain_lookup(host, live) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (host) params.append("host", host);
      if (live) params.append("live", live);
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

export function createGetDomainLookupTool() {
  return {
    definition: {
      name: 'get-domain-lookup',
      description: 'Domain Lookup',
      inputSchema: {
        type: 'object',
        properties: {
          host: {
            type: 'string',
            description: 'A domain name, hostname, FQDN, URL, HTML link or email address to lookup'
          },
          live: {
            type: 'boolean',
            description: 'For domains that we have never seen before then perform various live checks and realtime reconnaissance. <br>NOTE: this option may add additional non-deterministic delay to the request, if you require consistently fast API response times or just want to check our domain blocklists then you can disable this option'
          }
        },
        required: ["host"]
      }
    },
    handler: get_domain_lookup
  };
}