# Introduction
This repo contains a simple web app which takes a request and display a cowsay.

This [doc](https://hackmd.io/-YoGO4NrQaioK0W5HQXp8Q) describes how to deploy the app into a kind cluster.

## How to use BuildPack to build this app into an image

```shell
cd $GOPATH/src
git clone https://github.com/danniel1205/sample-cowsay-web-app.git
pack build --builder gcr.io/paketo-buildpacks/builder:tiny danielguo/cowsay-web-app:latest

docker run -p 1323:1323 danielguo/cowsay-web-app:latest
```

[`gcr.io/paketo-buildpacks/builder:tiny`](https://github.com/paketo-buildpacks/go#the-go-buildpack-is-compatible-with-the-following-builders)
is existing builder created for Go applications. This [notes](https://hackmd.io/-YoGO4NrQaioK0W5HQXp8Q) has the steps on
how I build this application and deploy on a K8S KIND cluster.

## How to use Carvel tools to build this app into an image 

TBA
