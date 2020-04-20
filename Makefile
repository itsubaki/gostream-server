install:
	set -x
	GO111MODULE=on go mod tidy
	GO111MODULE=on go install

.PHONY:
