FROM golang:1.15.5 AS builder
WORKDIR /go/src/github.com/frouioui/tagenal/api/users
COPY go.mod go.sum ./
COPY server/go.mod server/go.sum ./server/
COPY pb/go.mod pb/go.sum ./pb/
COPY db/go.mod db/go.sum ./db/
RUN go mod download
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/frouioui/tagenal/api/users/main /main
EXPOSE 10000 9090
ENTRYPOINT [ "/main" ]