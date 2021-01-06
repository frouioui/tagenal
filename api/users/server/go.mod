module github.com/frouioui/tagenal/api/users/server

go 1.15

replace github.com/frouioui/tagenal/api/users/pb => ../pb

replace github.com/frouioui/tagenal/api/users/db => ../db

require (
	github.com/frouioui/tagenal/api/users/db v0.0.0-00010101000000-000000000000
	github.com/frouioui/tagenal/api/users/pb v0.0.0-00010101000000-000000000000
	github.com/go-redis/cache/v8 v8.2.1
	github.com/go-redis/redis/extra/redisotel v0.2.0
	github.com/go-redis/redis/v8 v8.4.2
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/opentracing-contrib/go-grpc v0.0.0-20180928155321-4b5a12d3ff02
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.33.2
	vitess.io/vitess v0.7.0
)
