/**
 * UA Lookup
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

export async function get_ua_lookup(ua, ua-version, ua-platform, ua-platform-version, ua-mobile, device-model, device-brand) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (ua) params.append("ua", ua);
      if (ua-version) params.append("ua-version", ua-version);
      if (ua-platform) params.append("ua-platform", ua-platform);
      if (ua-platform-version) params.append("ua-platform-version", ua-platform-version);
      if (ua-mobile) params.append("ua-mobile", ua-mobile);
      if (device-model) params.append("device-model", device-model);
      if (device-brand) params.append("device-brand", device-brand);
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

export function createGetUaLookupTool() {
  return {
    definition: {
      name: 'get-ua-lookup',
      description: 'UA Lookup',
      inputSchema: {
        type: 'object',
        properties: {
          ua: {
            type: 'string',
            description: 'The user-agent string to lookup. For client hints use the 'UA' header or the JSON data directly from 'navigator.userAgentData.brands' or 'navigator.userAgentData.getHighEntropyValues()''
          },
          ua-version: {
            type: 'string',
            description: 'For client hints this corresponds to the 'UA-Full-Version' header or 'uaFullVersion' from NavigatorUAData'
          },
          ua-platform: {
            type: 'string',
            description: 'For client hints this corresponds to the 'UA-Platform' header or 'platform' from NavigatorUAData'
          },
          ua-platform-version: {
            type: 'string',
            description: 'For client hints this corresponds to the 'UA-Platform-Version' header or 'platformVersion' from NavigatorUAData'
          },
          ua-mobile: {
            type: 'string',
            description: 'For client hints this corresponds to the 'UA-Mobile' header or 'mobile' from NavigatorUAData'
          },
          device-model: {
            type: 'string',
            description: 'For client hints this corresponds to the 'UA-Model' header or 'model' from NavigatorUAData. <br>You can also use this parameter to lookup a device directly by its model name, model code or hardware code, on android you can get the model name from: https://developer.android.com/reference/android/os/Build.html#MODEL'
          },
          device-brand: {
            type: 'string',
            description: 'This parameter is only used in combination with 'device-model' when doing direct device lookups without any user-agent data. Set this to the brand or manufacturer name, this is required for accurate device detection with ambiguous model names. On android you can get the device brand from: https://developer.android.com/reference/android/os/Build#MANUFACTURER'
          }
        },
        required: ["ua"]
      }
    },
    handler: get_ua_lookup
  };
}