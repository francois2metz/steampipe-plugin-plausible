package plausible

import (
	"context"

	"github.com/francois2metz/go-plausible/plausible"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tablePlausibleTimeseries(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "plausible_timeseries",
		Description: "Provides timeseries data over a certain time period.",
		List: &plugin.ListConfig{
			Hydrate: listTimeseries,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "domain"},
				{Name: "period", Require: plugin.Optional},
				{Name: "date", Require: plugin.Optional},
				{Name: "interval", Require: plugin.Optional},
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
				Name:        "time",
				Type:        proto.ColumnType_STRING,
				Description: "The date of the serie.",
				Transform:   transform.FromField("Date"),
			},
			{
				Name:        "bounce_rate",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The bounce rate percentage.",
				Transform:   transform.FromMethod("BounceRate"),
			},
			{
				Name:        "pageviews",
				Type:        proto.ColumnType_INT,
				Description: "The number of pageview events.",
			},
			{
				Name:        "visit_duration",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The visit duration in seconds.",
				Transform:   transform.FromMethod("VisitDuration"),
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
			{
				Name:        "interval",
				Type:        proto.ColumnType_STRING,
				Description: "The reporting interval (date or month).",
				Transform:   transform.FromQual("interval"),
			},
		},
	}
}

func listTimeseries(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("plausible_timeseries.listTimeseries", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	domain := quals["domain"].GetStringValue()
	period := quals["period"].GetStringValue()
	date := quals["date"].GetStringValue()
	interval := quals["interval"].GetStringValue()
	if period == "" {
		period = "30d"
	}
	if interval == "" {
		interval = "day"
	}
	metrics := getMetrics(d.QueryContext.Columns)

	site := client.Site(domain)
	tsQuery := plausible.TimeseriesQuery{
		Period:   plausible.TimePeriod{Period: period, Date: date},
		Metrics:  metrics,
		Interval: plausible.TimeInterval(interval),
	}
	result, err := site.Timeseries(tsQuery)
	if err != nil {
		plugin.Logger(ctx).Error("plausible_timeseries.listTimeseries", err)
		return nil, err
	}
	for _, stat := range result {
		d.StreamListItem(ctx, stat)
	}
	return nil, nil
}
