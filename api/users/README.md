# Users API

The users api currently supports both gRPC and HTTP protocols.

## gRPC routes

The gRPC protobuf file can be found at `./pb/users.proto`.

## HTTP routes

Here is a list of all the implemented HTTP endpoints:

Route name | URL
--- | ---
Service Information | `/`
User By ID | `/id/{id}`
Users By Region | `/region/{region}`
New User | `/new`
New Bulk Users | `/new/bulk`