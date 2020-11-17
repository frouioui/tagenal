module github.com/frouioui/tagenal/mysql/api/users

go 1.15

require (
	github.com/frouioui/tagenal/mysql/api/users/pb v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	google.golang.org/grpc v1.33.2
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/frouioui/tagenal/mysql/api/users/pb => ./pb
