/**
 * Phone Validate
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

export async function get_phone_validate(number, country-code, ip) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (number) params.append("number", number);
      if (country-code) params.append("country-code", country-code);
      if (ip) params.append("ip", ip);
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

export function createGetPhoneValidateTool() {
  return {
    definition: {
      name: 'get-phone-validate',
      description: 'Phone Validate',
      inputSchema: {
        type: 'object',
        properties: {
          number: {
            type: 'string',
            description: 'A phone number. This can be in international format (E.164) or local format. If passing local format you must also set either the 'country-code' OR 'ip' options as well'
          },
          country-code: {
            type: 'string',
            description: 'ISO 2-letter country code, assume numbers are based in this country. If not set numbers are assumed to be in international format (with or without the leading + sign)'
          },
          ip: {
            type: 'string',
            description: 'Pass in a users IP address and we will assume numbers are based in the country of the IP address'
          }
        },
        required: ["number"]
      }
    },
    handler: get_phone_validate
  };
}