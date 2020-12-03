module github.com/frouioui/tagenal/api/articles/db

go 1.15

require (
	github.com/frouioui/tagenal/api/articles/pb v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/opentracing/opentracing-go v1.1.0
	google.golang.org/grpc v1.33.2
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.7
	vitess.io/vitess v0.7.0
)

replace github.com/frouioui/tagenal/api/articles/pb => ../pb
