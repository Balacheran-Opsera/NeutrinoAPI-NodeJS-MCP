/**
 * IP Blocklist Download
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

export async function get_ip_blocklist_download(format, include-vpn, cidr, ip6) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (format) params.append("format", format);
      if (include-vpn) params.append("include-vpn", include-vpn);
      if (cidr) params.append("cidr", cidr);
      if (ip6) params.append("ip6", ip6);
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

export function createGetIpBlocklistDownloadTool() {
  return {
    definition: {
      name: 'get-ip-blocklist-download',
      description: 'IP Blocklist Download',
      inputSchema: {
        type: 'object',
        properties: {
          format: {
            type: 'string',
            description: 'The data format. Can be either CSV or TXT'
          },
          include-vpn: {
            type: 'boolean',
            description: 'Include public VPN provider addresses, this option is only available for Tier 3 or higher accounts. Adds any IPs which are solely listed as VPN providers, IPs that are listed on multiple sensors will still be included without enabling this option. <br><b>WARNING</b>: This adds at least an additional 8 million IP addresses to the download if not using CIDR notation'
          },
          cidr: {
            type: 'boolean',
            description: 'Output IPs using CIDR notation. This option should be preferred but is off by default for backwards compatibility'
          },
          ip6: {
            type: 'boolean',
            description: 'Output the IPv6 version of the blocklist, the default is to output IPv4 only. Note that this option enables CIDR notation too as this is the only notation currently supported for IPv6'
          }
        },
        required: []
      }
    },
    handler: get_ip_blocklist_download
  };
}