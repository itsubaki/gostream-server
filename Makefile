install:
	set -x
	-rm ${GOPATH}/bin/gostream-api
	go install github.com/itsubaki/gostream-api

.PHONY:
