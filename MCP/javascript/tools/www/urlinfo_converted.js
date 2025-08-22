/**
 * URL Info
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

export async function get_url_info(url, timeout, retry, fetch-content, ignore-certificate-errors) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (url) params.append("url", url);
      if (timeout) params.append("timeout", timeout);
      if (retry) params.append("retry", retry);
      if (fetch-content) params.append("fetch-content", fetch-content);
      if (ignore-certificate-errors) params.append("ignore-certificate-errors", ignore-certificate-errors);
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

export function createGetUrlInfoTool() {
  return {
    definition: {
      name: 'get-url-info',
      description: 'URL Info',
      inputSchema: {
        type: 'object',
        properties: {
          url: {
            type: 'string',
            description: 'The URL to probe'
          },
          timeout: {
            type: 'number',
            description: 'Timeout in seconds. Give up if still trying to load the URL after this number of seconds'
          },
          retry: {
            type: 'number',
            description: 'If the request fails for any reason try again this many times'
          },
          fetch-content: {
            type: 'boolean',
            description: 'If this URL responds with html, text, json or xml then return the response. This option is useful if you want to perform further processing on the URL content (e.g. with the HTML Extract or HTML Clean APIs)'
          },
          ignore-certificate-errors: {
            type: 'boolean',
            description: 'Ignore any TLS/SSL certificate errors and load the URL anyway'
          }
        },
        required: ["url"]
      }
    },
    handler: get_url_info
  };
}