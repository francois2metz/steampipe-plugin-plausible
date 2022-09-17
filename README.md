# Plausible plugin for Steampipe

Use SQL to query statistics from [Plausible][].

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install francois2metz/plausible

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/plausible.spc ~/.steampipe/config/plausible.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[plausible]: https://plausible.io
