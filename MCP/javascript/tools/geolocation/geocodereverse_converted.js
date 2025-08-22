/**
 * Geocode Reverse
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

export async function get_geocode_reverse(latitude, longitude, language-code, zoom) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (latitude) params.append("latitude", latitude);
      if (longitude) params.append("longitude", longitude);
      if (language-code) params.append("language-code", language-code);
      if (zoom) params.append("zoom", zoom);
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

export function createGetGeocodeReverseTool() {
  return {
    definition: {
      name: 'get-geocode-reverse',
      description: 'Geocode Reverse',
      inputSchema: {
        type: 'object',
        properties: {
          latitude: {
            type: 'string',
            description: 'The location latitude in decimal degrees format'
          },
          longitude: {
            type: 'string',
            description: 'The location longitude in decimal degrees format'
          },
          language-code: {
            type: 'string',
            description: 'The language to display results in, available languages are: <ul> <li>de, en, es, fr, it, pt, ru</li> </ul>'
          },
          zoom: {
            type: 'string',
            description: 'The zoom level to respond with: <br> <ul> <li>address - the most precise address available</li> <li>street - the street level</li> <li>city - the city level</li> <li>state - the state level</li> <li>country - the country level</li> </ul>'
          }
        },
        required: ["latitude", "longitude"]
      }
    },
    handler: get_geocode_reverse
  };
}