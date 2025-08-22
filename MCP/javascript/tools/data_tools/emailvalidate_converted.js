/**
 * Email Validate
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

export async function get_email_validate(email, fix-typos) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (email) params.append("email", email);
      if (fix-typos) params.append("fix-typos", fix-typos);
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

export function createGetEmailValidateTool() {
  return {
    definition: {
      name: 'get-email-validate',
      description: 'Email Validate',
      inputSchema: {
        type: 'object',
        properties: {
          email: {
            type: 'string',
            description: 'An email address'
          },
          fix-typos: {
            type: 'boolean',
            description: 'Automatically attempt to fix typos in the address'
          }
        },
        required: ["email"]
      }
    },
    handler: get_email_validate
  };
}