module github.com/wzslr321/road_runner/server/users

go 1.19


replace github.com/gocql/gocql => github.com/scylladb/gocql v1.10.0

require (
	buf.build/gen/go/viago/auth/protocolbuffers/go v1.30.0-20230415224332-f5e8638609cf.1
	github.com/gocql/gocql v1.3.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/labstack/echo/v4 v4.10.2
	github.com/prometheus/client_golang v1.14.0
	github.com/scylladb/gocqlx/v2 v2.8.0
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.7.0
	google.golang.org/grpc v1.53.0
	buf.build/gen/go/viago/auth/bufbuild/connect-go v1.6.0-20230415224332-f5e8638609cf.1 // indirect
	buf.build/gen/go/viago/auth/grpc/go v1.3.0-20230415224332-f5e8638609cf.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bufbuild/connect-go v1.6.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)
