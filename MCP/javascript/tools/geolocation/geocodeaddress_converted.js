/**
 * Geocode Address
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

export async function get_geocode_address(address, house-number, street, city, county, state, postal-code, country-code, language-code, fuzzy-search) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (address) params.append("address", address);
      if (house-number) params.append("house-number", house-number);
      if (street) params.append("street", street);
      if (city) params.append("city", city);
      if (county) params.append("county", county);
      if (state) params.append("state", state);
      if (postal-code) params.append("postal-code", postal-code);
      if (country-code) params.append("country-code", country-code);
      if (language-code) params.append("language-code", language-code);
      if (fuzzy-search) params.append("fuzzy-search", fuzzy-search);
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

export function createGetGeocodeAddressTool() {
  return {
    definition: {
      name: 'get-geocode-address',
      description: 'Geocode Address',
      inputSchema: {
        type: 'object',
        properties: {
          address: {
            type: 'string',
            description: 'The full address, partial address or name of a place to try and locate. Comma separated address components are preferred.'
          },
          house-number: {
            type: 'string',
            description: 'The house/building number to locate'
          },
          street: {
            type: 'string',
            description: 'The street/road name to locate'
          },
          city: {
            type: 'string',
            description: 'The city/town name to locate'
          },
          county: {
            type: 'string',
            description: 'The county/region name to locate'
          },
          state: {
            type: 'string',
            description: 'The state name to locate'
          },
          postal-code: {
            type: 'string',
            description: 'The postal code to locate'
          },
          country-code: {
            type: 'string',
            description: 'Limit result to this country (the default is no country bias)'
          },
          language-code: {
            type: 'string',
            description: 'The language to display results in, available languages are: <ul> <li>de, en, es, fr, it, pt, ru, zh</li> </ul>'
          },
          fuzzy-search: {
            type: 'boolean',
            description: 'If no matches are found for the given address, start performing a recursive fuzzy search until a geolocation is found. This option is recommended for processing user input or implementing auto-complete. We use a combination of approximate string matching and data cleansing to find possible location matches'
          }
        },
        required: []
      }
    },
    handler: get_geocode_address
  };
}