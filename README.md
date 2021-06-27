# gostream-server

Stream Processing Server written in golang

## install

```console
$ go install github.com/itsubaki/gostream-server@latest
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
  "Time": "2021-06-27T01:59:21.781038+09:00",
  "Underlying": {
    "ID": "da6f6678-d69f-11eb-bba1-367dda957e1c",
    "Time": "2017-12-25T12:29:27Z",
    "Level": 4,
    "Message": "foobar"
  },
  "Record": {
    "count(*)": 1
  }
}
```
