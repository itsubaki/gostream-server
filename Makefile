install:
	set -x
	-rm ${GOPATH}/bin/gostream
	go install github.com/itsubaki/gostream-api

.PHONY:
