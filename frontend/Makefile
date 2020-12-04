export GO111MODULE=on

all: dockerbuild dockerpush

dockerpush:
	docker push frouioui/tagenal-frontend-service:latest

dockerbuild:
	docker build . -t frouioui/tagenal-frontend-service:latest

deps:
	go get -d -v ./...
