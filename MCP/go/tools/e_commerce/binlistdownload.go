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

func BinlistdownloadHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["include-iso3"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include-iso3=%v", val))
		}
		if val, ok := args["include-8digit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include-8digit=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/bin-list-download%s", cfg.BaseURL, queryString)
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
		var result string
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

func CreateBinlistdownloadTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_bin-list-download",
		mcp.WithDescription("BIN List Download"),
		mcp.WithBoolean("include-iso3", mcp.Description("Include ISO 3-letter country codes and ISO 3-letter currency codes in the data. These will be added to columns 10 and 11 respectively")),
		mcp.WithBoolean("include-8digit", mcp.Description("Include 8-digit and higher BIN codes. This option includes all 6-digit BINs and all 8-digit and higher BINs (including some 9, 10 and 11 digit BINs where available)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    BinlistdownloadHandler(cfg),
	}
}
