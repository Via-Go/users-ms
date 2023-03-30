FROM golang:1.20.1-alpine

WORKDIR /server/users

COPY src/go.mod ./
COPY src/go.sum ./

RUN go install github.com/cosmtrek/air@latest && \
    go mod download && \
    go mod verify

COPY src/ .

CMD air