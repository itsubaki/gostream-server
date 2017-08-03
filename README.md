# gostream
The Stream Processing Service written in Go

## install

```console
$ go get github.com/itsubaki/gostream
$ gostream
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /                         --> main.(Handler).POST-fm (3 handlers)
[GIN-debug] GET    /                         --> main.(Handler).GET-fm (3 handlers)
[GIN-debug] Listening and serving HTTP on :1234
...
```

## k8s

```console
$ sh ./sed.sh ${YOUR_DOCKER_IMAGE}
$ kubectl create -f kube/gostream-deployment.yml
$ kubectl create -f kube/gostream-service.yml
$ kubectl create -f kube/gostream-ingress.yml
```
