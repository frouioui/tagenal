module github.com/frouioui/tagenal/api/users/db

go 1.15

require (
	github.com/frouioui/tagenal/api/users/pb v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/opentracing/opentracing-go v1.1.0
	google.golang.org/grpc v1.53.0
	vitess.io/vitess v0.7.0
)

replace github.com/frouioui/tagenal/api/users/pb => ../pb
