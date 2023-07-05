module github.com/frouioui/tagenal/frontend

go 1.15

require (
	github.com/golang/protobuf v1.5.2
	github.com/labstack/echo-contrib v0.9.0
	github.com/labstack/echo/v4 v4.1.6
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/opentracing-contrib/go-grpc v0.0.0-20200813121455-4a6760c71486
	github.com/opentracing/opentracing-go v1.1.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

replace github.com/frouioui/tagenal/frontend/routes => ./routes
