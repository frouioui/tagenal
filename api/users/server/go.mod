module github.com/frouioui/tagenal/api/users/server

go 1.15

replace github.com/frouioui/tagenal/api/users/pb => ../pb
replace github.com/frouioui/tagenal/api/users/db => ../db

require (
	github.com/frouioui/tagenal/api/users/db v0.0.0-00010101000000-000000000000
	github.com/frouioui/tagenal/api/users/pb v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gorilla/mux v1.8.0
	google.golang.org/grpc v1.33.2
)