grafana-annotation: Post graphite annotation to grafana

# Install

go get -d "github.com/contentsquare/grafana-annotation"

# Configure

Will work with a configuration file (default to `~/.grafana-anotation-poster.yml`)

## Configuration file

```yaml
grafanaUri: https://some-grafana-host.tld
bearerToken: BearerTokenFromGrafana
```

## Create a Bearer Token

[Read the Docs](http://docs.grafana.org/http_api/auth/)

# Build

```
go build
```

# Call

## Options

```
Usage of grafana-annotation:
  -config-file string
    	Configuration File (default "~/.grafana-anotation-poster.yml")
  -data string
    	Additional data.
  -tag value
    	Tags. may be repeated multiple times
  -verbose
    	Be Verbose.
  -what string
    	The What item to post. (default "$(hostname)")
```

## Example call

```
~$ grafana-annotation -data "Details on this event" -tag foo \
  -tag bar -what "Something happened on system foo with bar event"
```
   
