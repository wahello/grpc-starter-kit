module github.com/xmlking/grpc-starter-kit

go 1.15

//replace github.com/xmlking/toolkit => /Users/schintha/Developer/Work/go/toolkit
replace github.com/xmlking/toolkit => github.com/xmlking/toolkit v0.1.2-0.20210110180932-82c7b80b90a9

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.15.0
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v0.15.0
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/cockroachdb/errors v1.8.2
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/facebook/ent v0.5.4
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.5
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/markbates/pkger v0.17.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/rs/zerolog v1.20.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/sercand/kuberesolver v2.4.0+incompatible
	github.com/soheilhy/cmux v0.1.4
	github.com/stretchr/testify v1.7.0
	github.com/tcfw/go-grpc-k8s-resolver v0.0.0-20201027075059-d3a2d14aa08f
	github.com/thoas/go-funk v0.7.0
	github.com/xmlking/toolkit v0.1.1
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.16.0
	go.opentelemetry.io/otel v0.16.0
	go.opentelemetry.io/otel/exporters/stdout v0.16.0
	go.opentelemetry.io/otel/sdk v0.16.0
	google.golang.org/grpc v1.35.0
	google.golang.org/grpc/examples v0.0.0-20210112202341-d3ae124a07fc // indirect
	google.golang.org/protobuf v1.25.0
)
