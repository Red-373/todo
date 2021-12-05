FROM golang:1.17.4-alpine3.15 AS builder

RUN go version
RUN apk add git

COPY ./ /github.com/Red-373/todo
WORKDIR /github.com/Red-373/todo

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

#lightweight docker container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/Red-373/todo/.bin/app .
COPY --from=0 /github.com/Red-373/todo/pkg/config/ ./config/

EXPOSE 8000

CMD [ "./app"]