package plausible

import (
	"context"

	"github.com/andrerfcsantos/go-plausible/plausible"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tablePlausibleSite(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "plausible_site",
		Description: "Sites from your plausible account.",
		List: &plugin.ListConfig{
			Hydrate: listSite,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSite,
			KeyColumns: plugin.SingleColumn("domain"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "domain",
				Type:        proto.ColumnType_STRING,
				Description: "The domain name of the site.",
			},
			{
				Name:        "timezone",
				Type:        proto.ColumnType_STRING,
				Description: "The timezone of the site.",
			},
		},
	}
}

func listSite(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("plausible_site.listSite", "connection_error", err)
		return nil, err
	}

	opts := plausible.ListSitesRequest{Limit: 100}
	for {
		result, err := client.ListSites(opts)
		if err != nil {
			plugin.Logger(ctx).Error("plausible_site.listSite", err)
			return nil, err
		}
		for _, site := range result.Sites {
			d.StreamListItem(ctx, site)
		}
		if result.Meta.After == "" {
			break
		}
		opts.After = result.Meta.After
	}
	return nil, nil
}


func getSite(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("plausible_site.getSite", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	domain := quals["domain"].GetStringValue()
	site := client.Site(domain)
	site2, err := site.Get()
	if err != nil {
		plugin.Logger(ctx).Error("plausible_site.getSite", err)
		return nil, err
	}
	return site2, err
}
