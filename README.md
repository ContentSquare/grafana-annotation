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


