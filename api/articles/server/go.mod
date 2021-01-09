module github.com/frouioui/tagenal/api/articles/server

go 1.15

replace github.com/frouioui/tagenal/api/articles/pb => ../pb

replace github.com/frouioui/tagenal/api/articles/db => ../db

require (
	github.com/frouioui/tagenal/api/articles/db v0.0.0-00010101000000-000000000000
	github.com/frouioui/tagenal/api/articles/pb v0.0.0-00010101000000-000000000000
	github.com/go-redis/cache/v8 v8.2.1
	github.com/go-redis/redis/extra/redisotel v0.2.0
	github.com/go-redis/redis/v8 v8.4.2
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/opentracing-contrib/go-grpc v0.0.0-20180928155321-4b5a12d3ff02
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible
	go.opentelemetry.io/otel v0.14.0
	go.opentelemetry.io/otel/exporters/stdout v0.14.0
	go.opentelemetry.io/otel/sdk v0.14.0
	google.golang.org/grpc v1.33.2
	gopkg.in/DataDog/dd-trace-go.v1 v1.17.0
)
