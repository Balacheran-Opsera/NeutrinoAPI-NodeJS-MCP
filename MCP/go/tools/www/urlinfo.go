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

func UrlinfoHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("url=%v", val))
		}
		if val, ok := args["fetch-content"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fetch-content=%v", val))
		}
		if val, ok := args["ignore-certificate-errors"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ignore-certificate-errors=%v", val))
		}
		if val, ok := args["timeout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeout=%v", val))
		}
		if val, ok := args["retry"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("retry=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/url-info%s", cfg.BaseURL, queryString)
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
		var result models.URLInfoResponse
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

func CreateUrlinfoTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_url-info",
		mcp.WithDescription("URL Info"),
		mcp.WithString("url", mcp.Required(), mcp.Description("The URL to probe")),
		mcp.WithBoolean("fetch-content", mcp.Description("If this URL responds with html, text, json or xml then return the response. This option is useful if you want to perform further processing on the URL content (e.g. with the HTML Extract or HTML Clean APIs)")),
		mcp.WithBoolean("ignore-certificate-errors", mcp.Description("Ignore any TLS/SSL certificate errors and load the URL anyway")),
		mcp.WithNumber("timeout", mcp.Description("Timeout in seconds. Give up if still trying to load the URL after this number of seconds")),
		mcp.WithNumber("retry", mcp.Description("If the request fails for any reason try again this many times")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UrlinfoHandler(cfg),
	}
}
