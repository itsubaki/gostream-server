# gostream-server

Stream Processing Server written in golang

## install

```console
$ go get github.com/itsubaki/gostream-server
```

## Example

```console
$ GOSTREAM_CONFIG=./config.yml gostream-server
config: {
  "Port":":1234",
  "Router":
    [
      {
        "Plugin":"LogEventPlugin",
        "Path":"/",
        "Query":"select count(*) from LogEvent.time(10 sec)"
      }
    ]
  }
```

```console
$ curl -s -X POST localhost:1234 -d '{"time":"2017-12-25T12:29:27Z", "Level": 4, "Message":"foobar"}' | jq .
{
  "ID":"2651d818-2c08-4895-981e-ddbf8e2614f8"
}
```

```console
$ curl -s localhost:1234 | jq .
{
  "ID":"2651d818-2c08-4895-981e-ddbf8e2614f8",
  "Level":4,
  "Message":"foobar",
  "Time":"2017-12-25T12:29:27Z",
  "count(*)":1
}
```
