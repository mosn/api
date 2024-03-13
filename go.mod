module mosn.io/api

go 1.17

require (
	// match version of envoy control plane.
	// See https://github.com/envoyproxy/go-control-plane/blob/v0.10.0/go.mod
	github.com/golang/protobuf v1.5.0
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0

)

require google.golang.org/protobuf v1.33.0 // indirect
