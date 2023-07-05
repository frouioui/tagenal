module github.com/frouioui/tagenal/api/users

go 1.15

require (
	github.com/frouioui/tagenal/api/users/server v0.0.0-00010101000000-000000000000
	github.com/opentracing-contrib/go-grpc v0.0.0-20200813121455-4a6760c71486 // indirect
	google.golang.org/grpc v1.53.0 // indirect
)

replace github.com/frouioui/tagenal/api/users/pb => ./pb

replace github.com/frouioui/tagenal/api/users/db => ./db

replace github.com/frouioui/tagenal/api/users/server => ./server
