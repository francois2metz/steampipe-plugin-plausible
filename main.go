package main

import (
	"github.com/francois2metz/steampipe-plugin-plausible/plausible"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: plausible.Plugin})
}
