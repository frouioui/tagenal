# Articles API

The articles api currently supports both gRPC and HTTP protocols.

## gRPC routes

The gRPC protobuf file can be found at `./pb/articles.proto`.

## HTTP routes

Here is a list of all the implemented HTTP endpoints:

Route name | URL
--- | ---
Service Information | `/`
Article By ID | `/id/{id}`
Articles By Category | `/category/{category}`
Articles Stored In Region | `/region/id/{region_id}`
New Article | `/new`
New Bulk Articles | `/new/bulk`