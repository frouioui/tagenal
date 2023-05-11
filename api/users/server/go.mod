module github.com/frouioui/tagenal/api/users/server

go 1.15

replace github.com/frouioui/tagenal/api/users/pb => ../pb

replace github.com/frouioui/tagenal/api/users/db => ../db

require (
	github.com/frouioui/tagenal/api/users/db v0.0.0-00010101000000-000000000000
	github.com/frouioui/tagenal/api/users/pb v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/extra/redisotel v0.2.0
	github.com/go-redis/redis/v8 v8.4.2
	github.com/gorilla/mux v1.8.0
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	google.golang.org/grpc v1.52.3
	vitess.io/vitess v0.16.2 // indirect
)
