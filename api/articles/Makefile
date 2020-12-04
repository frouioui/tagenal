export GO111MODULE=on

all: protobuild dockerbuild dockerpush

dockerpush:
	docker push frouioui/tagenal-articles-service:latest

dockerbuild:
	docker build . -t frouioui/tagenal-articles-service:latest

protobuild:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pb/articles.proto

build:
	go build main.go

clean:
	go clean

deps:
	go get -d -v ./...
	go get -d -v google.golang.org/grpc/...