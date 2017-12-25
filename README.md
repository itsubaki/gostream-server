# gostream
The Stream Processing Service written in Go

## install

```console
$ go get github.com/itsubaki/gostream
```

## Example

```console
export GOSTREAM_LISTEN_PORT=1234
$ ./gostream
config: {"Port":":1234"}
2017-12-25 15:40:44.608425401 +0900 JST
```

 - Publish

```console
$ curl -X POST localhost:1234 -d '{"time":"2017-12-25T12:29:27Z", "Level": 4, "Message":"foobar"}'
{ "ok" }
```

 - Subscribe

```console
$ curl localhost:1234
{"Time":"2017-12-25T15:40:49.949279107+09:00","Underlying":{"time":"2017-12-25T12:29:27Z","level":4,"message":"foobar"},"Record":{"count":1}}
```
