---
organization: francois2metz
category: ["saas"]
brand_color: "#4052c7"
display_name: "Plausible"
short_name: "plausible"
description: "Steampipe plugin for querying statistics from Plausible."
og_description: "Query Plausible with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/francois2metz/plausible-social-graphic.png"
icon_url: "/images/plugins/francois2metz/plausible.svg"
---

# Plausible + Steampipe

[Plausible](https://plausible.io/) is an analytics company.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
```

```
```

## Documentation

- **[Table definitions & examples →](/plugins/francois2metz/plausible/tables)**

## Get started

### Install

Download and install the latest Plausible plugin:

```bash
steampipe plugin install francois2metz/plausible
```

### Configuration

Installing the latest plausible plugin will create a config file (`~/.steampipe/config/plausible.spc`) with a single connection named `plausible`:

```hcl
connection "francois2metz/plausible" {
  plugin = "francois2metz/plausible"

  # Go to https://plausible.io/settings/api-keys/new
  # token = ""
}

```

You can also use environment variables:

- `PLAUSIBLE_TOKEN`: Your plausible API Key

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-plausible
