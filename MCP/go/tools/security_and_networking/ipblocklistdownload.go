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

func IpblocklistdownloadHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["format"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("format=%v", val))
		}
		if val, ok := args["include-vpn"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include-vpn=%v", val))
		}
		if val, ok := args["cidr"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cidr=%v", val))
		}
		if val, ok := args["ip6"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ip6=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/ip-blocklist-download%s", cfg.BaseURL, queryString)
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

func CreateIpblocklistdownloadTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_ip-blocklist-download",
		mcp.WithDescription("IP Blocklist Download"),
		mcp.WithString("format", mcp.Description("The data format. Can be either CSV or TXT")),
		mcp.WithBoolean("include-vpn", mcp.Description("Include public VPN provider addresses, this option is only available for Tier 3 or higher accounts. Adds any IPs which are solely listed as VPN providers, IPs that are listed on multiple sensors will still be included without enabling this option. <br><b>WARNING</b>: This adds at least an additional 8 million IP addresses to the download if not using CIDR notation")),
		mcp.WithBoolean("cidr", mcp.Description("Output IPs using CIDR notation. This option should be preferred but is off by default for backwards compatibility")),
		mcp.WithBoolean("ip6", mcp.Description("Output the IPv6 version of the blocklist, the default is to output IPv4 only. Note that this option enables CIDR notation too as this is the only notation currently supported for IPv6")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    IpblocklistdownloadHandler(cfg),
	}
}
