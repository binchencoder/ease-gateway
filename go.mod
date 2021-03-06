module github.com/binchencoder/ease-gateway

go 1.13

require (
	github.com/binchencoder/gateway-proto v0.0.7
	github.com/binchencoder/letsgo v0.0.3
	github.com/binchencoder/skylb-api v0.0.5
	github.com/fatih/color v1.9.0
	github.com/ghodss/yaml v1.0.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1
	github.com/klauspost/compress v1.10.10
	github.com/pborman/uuid v1.2.0
	github.com/prometheus/client_golang v1.7.1
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/genproto v0.0.0-20201019141844-1ed22bb0c154
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
	google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.29.1
	google.golang.org/grpc v1.33.1 => google.golang.org/grpc v1.29.1
)
