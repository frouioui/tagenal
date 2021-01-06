module github.com/frouioui/tagenal/frontend

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo-contrib v0.9.0
	github.com/labstack/echo/v4 v4.1.6
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.19.1-0.20191002155754-0be28c34dabf+incompatible
	google.golang.org/grpc v1.33.2
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1 // indirect
	google.golang.org/protobuf v1.25.0
)

replace github.com/frouioui/tagenal/frontend/routes => ./routes
