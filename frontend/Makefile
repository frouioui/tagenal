export GO111MODULE=on

all: dockerbuild dockerpush

dockerpush:
	docker push frouioui/tagenal-frontend-service:latest

dockerbuild: clientproto
	docker build . -t frouioui/tagenal-frontend-service:latest

build:
	go build main.go

clientproto:
	cp ../api/articles/pb/*.pb.go ./client/pb/articles
	cp ../api/users/pb/*.pb.go ./client/pb/users

clean:
	go clean

deps:
	go get -d -v ./...
