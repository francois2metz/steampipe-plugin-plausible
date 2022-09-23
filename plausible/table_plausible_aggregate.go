package plausible

import (
	"context"
	"github.com/francois2metz/go-plausible/plausible"
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
				Description: "The bounce rate percentage.",
			},
			{
				Name:        "pageviews",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The number of pageview events.",
			},
			{
				Name:        "visit_duration",
				Type:        proto.ColumnType_INT,
				Description: "The visit duration in seconds.",
			},
			{
				Name:        "visitors",
				Type:        proto.ColumnType_INT,
				Description: "The number of unique visitors.",
			},
			{
				Name:        "visits",
				Type:        proto.ColumnType_INT,
				Description: "The number of visits/sessions.",
			},
			{
				Name:        "events",
				Type:        proto.ColumnType_INT,
				Description: "The number of events (pageviews + custom events).",
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
	metrics := plausible.Metrics{}
	for _, v := range d.QueryContext.Columns {
		switch v {
		case "bounce_rate":
			metrics = append(metrics, plausible.BounceRate)
		case "visitors":
			metrics = append(metrics, plausible.Visitors)
		case "pageviews":
			metrics = append(metrics, plausible.PageViews)
		case "visit_duration":
			metrics = append(metrics, plausible.VisitDuration)
		case "visits":
			metrics = append(metrics, plausible.Visits)
		case "events":
			metrics = append(metrics, plausible.Events)
		}
	}

	site := client.Site(domain)
	visitorsQuery := plausible.AggregateQuery{
		Period:  plausible.TimePeriod{Period: period, Date: date},
		Metrics: metrics,
	}
	result, err := site.Aggregate(visitorsQuery)
	if err != nil {
		plugin.Logger(ctx).Error("plausible_aggregate.getAggregate", err)
		return nil, err
	}
	return result, nil
}
