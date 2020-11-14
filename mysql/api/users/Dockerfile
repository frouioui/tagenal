FROM golang:1.15.5 AS builder

WORKDIR /go/src/github.com/frouioui/tagenal/mysql/api/users

COPY . /go/src/github.com/frouioui/tagenal/mysql/api/users

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# -------------------

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/frouioui/tagenal/mysql/api/users/main .

EXPOSE 8080

ENV PORT=8080

CMD [ "./main" ]

LABEL service="tagenal_users"