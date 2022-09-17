package plausible

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

type plausibleCurrentVisitor struct {
	Count int
}

func tablePlausibleCurrentVisitor(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "plausible_current_visitor",
		Description: "A current visitor is defined as a visitor who triggered a pageview on your site in the last 5 minutes.",
		Get: &plugin.GetConfig{
			Hydrate:    getCurrentVisitor,
			KeyColumns: plugin.SingleColumn("domain"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "domain",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the app.",
				Transform:   transform.FromQual("domain"),
			},
			{
				Name:        "count",
				Type:        proto.ColumnType_INT,
				Description: "Number of current visitors.",
			},
		},
	}
}

func getCurrentVisitor(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("plausible_current_visitor.getCurrentVisitor", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	domain := quals["domain"].GetStringValue()
	site := client.Site(domain)
	visitors, err := site.CurrentVisitors()
	if err != nil {
		plugin.Logger(ctx).Error("plausible_current_visitor.getCurrentVisitor", err)
		return nil, err
	}
	return plausibleCurrentVisitor{Count: visitors}, nil
}
