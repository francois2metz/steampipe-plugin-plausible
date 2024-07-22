package plausible

import (
	"context"
	"strings"

	"github.com/andrerfcsantos/go-plausible/plausible"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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
				Name:        "bounce_rate_change",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The percent difference of the bounce rate with the previous period.",
			},
			{
				Name:        "pageviews",
				Type:        proto.ColumnType_INT,
				Description: "The number of pageview events.",
			},
			{
				Name:        "pageviews_change",
				Type:        proto.ColumnType_INT,
				Description: "The percent difference of the number of pageview events with the previous period.",
			},
			{
				Name:        "visit_duration",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The visit duration in seconds.",
			},
			{
				Name:        "visit_duration_change",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The percent difference of the visit duration in seconds with the previous period.",
			},
			{
				Name:        "visitors",
				Type:        proto.ColumnType_INT,
				Description: "The number of unique visitors.",
			},
			{
				Name:        "visitors_change",
				Type:        proto.ColumnType_INT,
				Description: "The percent difference of the number of unique visitors with the previous period.",
			},
			{
				Name:        "visits",
				Type:        proto.ColumnType_INT,
				Description: "The number of visits/sessions.",
			},
			{
				Name:        "visits_change",
				Type:        proto.ColumnType_INT,
				Description: "The percent difference of the number of visits/sessions with the previous period.",
			},
			{
				Name:        "events",
				Type:        proto.ColumnType_INT,
				Description: "The number of events (pageviews + custom events).",
			},
			{
				Name:        "events_change",
				Type:        proto.ColumnType_INT,
				Description: "The percent difference of the number of events (pageviews + custom events) with the previous period.",
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
	quals := d.EqualsQuals
	domain := quals["domain"].GetStringValue()
	period := quals["period"].GetStringValue()
	date := quals["date"].GetStringValue()
	if period == "" {
		period = "30d"
	}
	metrics := getMetrics(d.QueryContext.Columns)
	comparePreviousPeriod := false
	for _, v := range d.QueryContext.Columns {
		if strings.Contains(v, "_change") {
			comparePreviousPeriod = true
		}
	}

	site := client.Site(domain)
	visitorsQuery := plausible.AggregateQuery{
		Period:                plausible.TimePeriod{Period: period, Date: date},
		Metrics:               metrics,
		ComparePreviousPeriod: comparePreviousPeriod,
	}
	result, err := site.Aggregate(visitorsQuery)
	if err != nil {
		plugin.Logger(ctx).Error("plausible_aggregate.getAggregate", err)
		return nil, err
	}
	return result, nil
}
