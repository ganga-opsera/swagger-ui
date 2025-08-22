package main

import (
	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	tools_password "github.com/mcp-server/mcp-server/tools/password"
	tools_application "github.com/mcp-server/mcp-server/tools/application"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_password.CreateGet_passwordTool(cfg),
		tools_application.CreateGet_applicationTool(cfg),
	}
}
