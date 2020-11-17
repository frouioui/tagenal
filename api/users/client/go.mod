module github.com/frouioui/tagenal/mysql/api/users/client

go 1.15

replace github.com/frouioui/tagenal/mysql/api/users/pb => ./pb

require (
	github.com/frouioui/tagenal/mysql/api/users/pb v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	google.golang.org/grpc v1.33.2
)
