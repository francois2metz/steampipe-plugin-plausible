package plausible

import (
	"context"
	"errors"
	"os"

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

	plausibleConfig := GetConfig(d.Connection)

	if plausibleConfig.Token != nil {
		token = *plausibleConfig.Token
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file or set the PLAUSIBLE_TOKEN environment variable and then restart Steampipe")
	}

	client := plausible.NewClient(token)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
