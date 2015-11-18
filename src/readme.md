# Go Server

The go server has three main parts:

* The core exporter
* A command-line interface
* A web interface using websockets


## Assets

All resources are encoded and stored within the app using go-bindata.

For development:

```go-bindata -o assets.go -debug data/...```

For production:

```go-bindata -o assets.go data/...```
