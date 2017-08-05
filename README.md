# gostream
The Stream Processing Service written in Go

## install

```console
$ go get -u cloud.google.com/go/...
$ go get github.com/itsubaki/gostream
$ gostream
[GIN-debug] Listening and serving HTTP on :1234
...
```

## k8s

```console
$ sh ./sed.sh ${YOUR_DOCKER_IMAGE}
$ kubectl create -f kube/deployment.yml.tmp
$ kubectl create -f kube/service.yml
$ kubectl create -f kube/ingress.yml
```

## Output

 - ```stdout```, ```logging```, ```pubsub```, ```spanner```
 - default is ```stdout```

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
$ export GOSTREAM_SPANNER_DATABASE=${YOUR_GCP_SPANNER_DATABASE}
$ exprot GOSTREAM_SPANNER_TABLE=${YOUR_GCP_SPANNER_TABLE}
$ exprot GOSTREAM_SPANNER_COLUMN=${YOUR_GCP_SPANNER_COLUMN}
```
