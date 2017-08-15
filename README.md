# gostream
The Stream Processing Service written in Go

## install

```console
$ go get github.com/itsubaki/gostream
```

## k8s

```console
$ sh ./sed.sh ${YOUR_DOCKER_IMAGE}
$ kubectl create -f gostream.yml.tmp
```

## Output

 - ```stdout```(default)
 - ```logging```, ```pubsub```, ```spanner```

```console
$ export GOOGLE_APPLICATION_CREDENTIALS=`pwd`/credential.json
$ export GOSTREAM_PROJECT_ID=${YOUR_GCP_PROJECT_ID}
```

### Logging

```console
$ export GOSTREAM_OUTPUT=logging
$ export GOSTREAM_LOGGING_LOGGER=${YOUR_GCP_LOGGING_LOGGER}
```

### PubSub

```console
$ export GOSTREAM_OUTPUT=pubsub
$ export GOSTREAM_PUBSUB_TOPIC=${YOUR_GCP_PUBSUB_TOPIC}
```

### Spanner

```console
$ export GOSTREAM_OUTPUT=spanner
$ export GOSTREAM_SPANNER=${YOUR_GCP_SPANNER}
```
