package plausible

import (
	"context"
	"github.com/andrerfcsantos/go-plausible/plausible"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tablePlausibleAggregate(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "plausible_aggregate",
		Description: "Aggregates metrics over a certain time period.",
		Get: &plugin.GetConfig{
			Hydrate:    getAggregate,
			KeyColumns: plugin.SingleColumn("domain"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "domain",
				Type:        proto.ColumnType_STRING,
				Description: "Domain name.",
				Transform:   transform.FromQual("domain"),
			},
			{
				Name:        "bounce_rate",
				Type:        proto.ColumnType_DOUBLE,
				Description: "",
			},
			{
				Name:        "pageviews",
				Type:        proto.ColumnType_DOUBLE,
				Description: "",
			},
			{
				Name:        "visit_duration",
				Type:        proto.ColumnType_DOUBLE,
				Description: "",
			},
			{
				Name:        "visitors",
				Type:        proto.ColumnType_DOUBLE,
				Description: "",
			},
		},
	}
}

func getAggregate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("plausible_aggregate.getAggregate", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	domain := quals["domain"].GetStringValue()
	site := client.Site(domain)
	visitorsQuery := plausible.AggregateQuery{
		Period: plausible.DayPeriod(),
		Metrics: plausible.Metrics{
			plausible.Visitors,
			plausible.PageViews,
			plausible.BounceRate,
			plausible.VisitDuration,
		},
	}
	result, err := site.Aggregate(visitorsQuery)
	if err != nil {
		plugin.Logger(ctx).Error("plausible_aggregate.getAggregate", err)
		return nil, err
	}
	return result, nil
}
