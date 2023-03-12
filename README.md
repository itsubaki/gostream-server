# gostream-server

Stream Processing Server written in Go

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
  "ID": "002d4f48-6d62-11ec-bf09-367dda957e1c"
}
```

```console
$ curl -s localhost:1234 | jq .
{
  "time": "2022-01-04T22:27:01.883604+09:00",
  "underlying": {
    "ID": "002d4f48-6d62-11ec-bf09-367dda957e1c",
    "Time": "2017-12-25T12:29:27Z",
    "Level": 4,
    "Message": "foobar"
  },
  "result_set": [
    1
  ]
}
```
