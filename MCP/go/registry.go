package main

import (
	"github.com/neutrino-api/mcp-server/config"
	"github.com/neutrino-api/mcp-server/models"
	tools_www "github.com/neutrino-api/mcp-server/tools/www"
	tools_security_and_networking "github.com/neutrino-api/mcp-server/tools/security_and_networking"
	tools_e_commerce "github.com/neutrino-api/mcp-server/tools/e_commerce"
	tools_data_tools "github.com/neutrino-api/mcp-server/tools/data_tools"
	tools_telephony "github.com/neutrino-api/mcp-server/tools/telephony"
	tools_geolocation "github.com/neutrino-api/mcp-server/tools/geolocation"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_www.CreateUrlinfoTool(cfg),
		tools_security_and_networking.CreateIpblocklistdownloadTool(cfg),
		tools_e_commerce.CreateBinlistdownloadTool(cfg),
		tools_data_tools.CreateEmailvalidateTool(cfg),
		tools_telephony.CreateHlrlookupTool(cfg),
		tools_geolocation.CreateGeocodereverseTool(cfg),
		tools_security_and_networking.CreateHostreputationTool(cfg),
		tools_telephony.CreateVerifysecuritycodeTool(cfg),
		tools_security_and_networking.CreateIpblocklistTool(cfg),
		tools_e_commerce.CreateBinlookupTool(cfg),
		tools_e_commerce.CreateConvertTool(cfg),
		tools_security_and_networking.CreateDomainlookupTool(cfg),
		tools_geolocation.CreateGeocodeaddressTool(cfg),
		tools_data_tools.CreatePhonevalidateTool(cfg),
		tools_geolocation.CreateIpinfoTool(cfg),
		tools_security_and_networking.CreateIpprobeTool(cfg),
		tools_security_and_networking.CreateEmailverifyTool(cfg),
		tools_data_tools.CreateUalookupTool(cfg),
	}
}
