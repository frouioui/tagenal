FROM golang:1.15.5 AS builder
WORKDIR /go/src/github.com/frouioui/tagenal/jobs/popularity
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/frouioui/tagenal/jobs/popularity/main /main