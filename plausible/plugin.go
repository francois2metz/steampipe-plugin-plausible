package plausible

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-plausible",
		DefaultTransform: transform.FromGo(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"plausible_site":            tablePlausibleSite(ctx),
			"plausible_aggregate":       tablePlausibleAggregate(ctx),
			"plausible_current_visitor": tablePlausibleCurrentVisitor(ctx),
			"plausible_timeseries":      tablePlausibleTimeseries(ctx),
		},
	}
	return p
}
