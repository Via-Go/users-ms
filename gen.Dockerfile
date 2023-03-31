FROM golang:1.20.1-alpine

WORKDIR /server/users

COPY go.mod ./
COPY go.sum ./

RUN go mod download && \
    go mod verify && \
    apk update && apk add --no-cache make protobuf-dev

COPY scrap_proto.go .
COPY src/proto-gen/ ./src/proto-gen/

CMD go run scrap_proto.go