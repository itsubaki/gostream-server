run:
	set -x
	GOSTREAM_CONFIG=./config.yml go run main.go

curl:
	curl -s -X POST localhost:1234 -d '{"time":"2017-12-25T12:29:27Z", "Level": 4, "Message":"foobar"}' | jq .
	curl -s localhost:1234 | jq .

.PHONY:
