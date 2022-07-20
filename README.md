# A prometheus exporter for Exoscale

## Help

```sh
NAME:
   exoscale_exporter - Fetches statistics from Exoscale API

USAGE:
   exoscale_exporter [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config.file value         Path to configuration file. (default: "config.yaml") [$CONFIG_FILE]
   --help, -h                  show help (default: false)
   --web.listen-address value  Address to listen on for web interface and telemetry. (default: ":9116") [$LISTEN_ADDRESS]

```

## Example config

**config.yaml**

```yaml
exoscale_config:
  key: 'EXO...'
  secret: '...'
```
