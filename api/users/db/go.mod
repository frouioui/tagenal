module github.com/frouioui/tagenal/api/users/db

go 1.15

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/frouioui/tagenal/api/users/pb v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.7.0
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/grpc v1.52.3
	vitess.io/vitess v0.16.2
)

replace github.com/frouioui/tagenal/api/users/pb => ../pb
