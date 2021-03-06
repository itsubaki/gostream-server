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
$ curl -X POST localhost:1234 -d '{"time":"2017-12-25T12:29:27Z", "Level": 4, "Message":"foobar"}'
{
  "ID":"2651d818-2c08-4895-981e-ddbf8e2614f8"
}
```

```console
$ curl localhost:1234
{
  "Time":"2018-05-16T15:08:02.993021138+09:00",
  "Underlying":
    {
      "ID":"2651d818-2c08-4895-981e-ddbf8e2614f8",
      "Time":"2017-12-25T12:29:27Z",
      "Level":4,
      "Message":"foobar"
    },
  "Record":
    {
      "ID":"2651d818-2c08-4895-981e-ddbf8e2614f8",
      "Level":4,
      "Message":"foobar",
      "Time":"2017-12-25T12:29:27Z",
      "count(*)":1
    }
}
```
