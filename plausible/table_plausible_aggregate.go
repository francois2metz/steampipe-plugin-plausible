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
			Hydrate: getAggregate,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "domain"},
				{Name: "period", Require: plugin.Optional},
				{Name: "date", Require: plugin.Optional},
			},
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
			{
				Name:        "period",
				Type:        proto.ColumnType_STRING,
				Description: "The time period (12mo, 6mo, month, 30d, 7d, day, custom). Default is 30d.",
				Transform:   transform.FromQual("period"),
			},
			{
				Name:        "date",
				Type:        proto.ColumnType_STRING,
				Description: "The date range to aggregate from when using the day or custom period.",
				Transform:   transform.FromQual("date"),
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
	period := quals["period"].GetStringValue()
	date := quals["date"].GetStringValue()
	if period == "" {
		period = "30d"
	}

	site := client.Site(domain)
	visitorsQuery := plausible.AggregateQuery{
		Period: plausible.TimePeriod{Period: period, Date: date},
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
