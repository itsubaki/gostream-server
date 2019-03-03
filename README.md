# gostream
The Stream Processing Service written in Go

## install

```console
$ go get github.com/itsubaki/gostream
```

## Example

```console
$ export GOSTREAM_CONFIG=./gostream.yml
$ gostream
config: {
  "Port":":1234",
  "Router":
    [
      {
        "Path":"/",
        "Query":"select count(*) from LogEvent.time(10 sec)"
      }
    ]
  }
```

 - Publish

```console
$ curl -X POST localhost:1234 -d '{"time":"2017-12-25T12:29:27Z", "Level": 4, "Message":"foobar"}'
{
  "ID":"2651d818-2c08-4895-981e-ddbf8e2614f8"
}
```

 - Subscribe

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
