module github.com/frouioui/tagenal/api/articles

go 1.15

require (
	github.com/frouioui/tagenal/api/articles/server v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.53.0 // indirect
)

replace github.com/frouioui/tagenal/api/articles/pb => ./pb

replace github.com/frouioui/tagenal/api/articles/db => ./db

replace github.com/frouioui/tagenal/api/articles/server => ./server
