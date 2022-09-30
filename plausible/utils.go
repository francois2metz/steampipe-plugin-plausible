package plausible

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/andrerfcsantos/go-plausible/plausible"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*plausible.Client, error) {
	// get plausible client from cache
	cacheKey := "plausible"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*plausible.Client), nil
	}

	token := os.Getenv("PLAUSIBLE_TOKEN")
	baseURL := os.Getenv("PLAUSIBLE_BASE_URL")

	plausibleConfig := GetConfig(d.Connection)

	if plausibleConfig.Token != nil {
		token = *plausibleConfig.Token
	}
	if plausibleConfig.BaseURL != nil {
		baseURL = *plausibleConfig.BaseURL
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file or set the PLAUSIBLE_TOKEN environment variable and then restart Steampipe")
	}

	var client *plausible.Client
	if baseURL != "" {
		client = plausible.NewClientWithBaseURL(token, baseURL)
	} else {
		client = plausible.NewClient(token)
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

func getMetrics(columns []string) plausible.Metrics {
	metrics := plausible.Metrics{}
	for _, v := range columns {
		switch strings.Replace(v, "_change", "", -1) {
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
	return metrics
}
