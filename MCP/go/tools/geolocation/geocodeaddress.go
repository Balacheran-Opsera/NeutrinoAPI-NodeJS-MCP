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

func GeocodeaddressHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["address"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("address=%v", val))
		}
		if val, ok := args["house-number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("house-number=%v", val))
		}
		if val, ok := args["street"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("street=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["county"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("county=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["postal-code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("postal-code=%v", val))
		}
		if val, ok := args["country-code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country-code=%v", val))
		}
		if val, ok := args["language-code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("language-code=%v", val))
		}
		if val, ok := args["fuzzy-search"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fuzzy-search=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/geocode-address%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("user-id", cfg.APIKey)
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
		var result models.GeocodeAddressResponse
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

func CreateGeocodeaddressTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_geocode-address",
		mcp.WithDescription("Geocode Address"),
		mcp.WithString("address", mcp.Description("The full address, partial address or name of a place to try and locate. Comma separated address components are preferred.")),
		mcp.WithString("house-number", mcp.Description("The house/building number to locate")),
		mcp.WithString("street", mcp.Description("The street/road name to locate")),
		mcp.WithString("city", mcp.Description("The city/town name to locate")),
		mcp.WithString("county", mcp.Description("The county/region name to locate")),
		mcp.WithString("state", mcp.Description("The state name to locate")),
		mcp.WithString("postal-code", mcp.Description("The postal code to locate")),
		mcp.WithString("country-code", mcp.Description("Limit result to this country (the default is no country bias)")),
		mcp.WithString("language-code", mcp.Description("The language to display results in, available languages are: <ul> <li>de, en, es, fr, it, pt, ru, zh</li> </ul>")),
		mcp.WithBoolean("fuzzy-search", mcp.Description("If no matches are found for the given address, start performing a recursive fuzzy search until a geolocation is found. This option is recommended for processing user input or implementing auto-complete. We use a combination of approximate string matching and data cleansing to find possible location matches")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GeocodeaddressHandler(cfg),
	}
}
