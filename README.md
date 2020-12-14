grafana-annotation: Post annotation to Elastic Search

# Install

go get -d "github.com/contentsquare/grafana-annotation"

# Build

```shell
make build
```

# Configure

Will work with a configuration file (default to `~/.grafana-anotation-poster.yml`)

## Configuration file

```yaml
region: eu-west-1
env: dev
provider: aws
role: testproject
elasticsearch:
  bootstrap_servers:
    - http://localhost:9200
  index_name: "test_index.2006.02"
```


## Example usage


```shell script
~# grafana-annotation --config-file resources/config-test.yml post --tags systemd --tags stop --tags kafka service kafka is stopped
Using config file: resources/config-test.yml
[2020-12-10 13:27:29]  INFO Annotation successfully posted to elasticsearch. tags=[systemd stop kafka eu-west-1 testproject dev aws Vianneys-MacBook-Pro.local]
~#
```

