install:
	set -x
	GO111MODULE=on go mod tidy
	GO111MODULE=on go install

run:
	set -x
	GOSTREAM_CONFIG=./config.yml gostream-server

.PHONY:
