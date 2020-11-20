module github.com/frouioui/tagenal/api/users/db

go 1.15

require (
	github.com/frouioui/tagenal/api/users/pb v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
)

replace github.com/frouioui/tagenal/api/users/pb => ../pb
