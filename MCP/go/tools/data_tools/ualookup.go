package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/neutrino-api/mcp-server/config"
	"github.com/neutrino-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UalookupHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["ua"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ua=%v", val))
		}
		if val, ok := args["ua-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ua-version=%v", val))
		}
		if val, ok := args["ua-platform"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ua-platform=%v", val))
		}
		if val, ok := args["ua-platform-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ua-platform-version=%v", val))
		}
		if val, ok := args["ua-mobile"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ua-mobile=%v", val))
		}
		if val, ok := args["device-model"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("device-model=%v", val))
		}
		if val, ok := args["device-brand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("device-brand=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/ua-lookup%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("api-key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.UALookupResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateUalookupTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_ua-lookup",
		mcp.WithDescription("UA Lookup"),
		mcp.WithString("ua", mcp.Required(), mcp.Description("The user-agent string to lookup. For client hints use the 'UA' header or the JSON data directly from 'navigator.userAgentData.brands' or 'navigator.userAgentData.getHighEntropyValues()'")),
		mcp.WithString("ua-version", mcp.Description("For client hints this corresponds to the 'UA-Full-Version' header or 'uaFullVersion' from NavigatorUAData")),
		mcp.WithString("ua-platform", mcp.Description("For client hints this corresponds to the 'UA-Platform' header or 'platform' from NavigatorUAData")),
		mcp.WithString("ua-platform-version", mcp.Description("For client hints this corresponds to the 'UA-Platform-Version' header or 'platformVersion' from NavigatorUAData")),
		mcp.WithString("ua-mobile", mcp.Description("For client hints this corresponds to the 'UA-Mobile' header or 'mobile' from NavigatorUAData")),
		mcp.WithString("device-model", mcp.Description("For client hints this corresponds to the 'UA-Model' header or 'model' from NavigatorUAData. <br>You can also use this parameter to lookup a device directly by its model name, model code or hardware code, on android you can get the model name from: https://developer.android.com/reference/android/os/Build.html#MODEL")),
		mcp.WithString("device-brand", mcp.Description("This parameter is only used in combination with 'device-model' when doing direct device lookups without any user-agent data. Set this to the brand or manufacturer name, this is required for accurate device detection with ambiguous model names. On android you can get the device brand from: https://developer.android.com/reference/android/os/Build#MANUFACTURER")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UalookupHandler(cfg),
	}
}
