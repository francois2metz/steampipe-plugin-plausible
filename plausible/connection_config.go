package plausible

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type plausibleConfig struct {
	Token *string `cty:"token"`
	BaseURL *string `cty:"base_url"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"base_url": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &plausibleConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) plausibleConfig {
	if connection == nil || connection.Config == nil {
		return plausibleConfig{}
	}
	config, _ := connection.Config.(plausibleConfig)
	return config
}
