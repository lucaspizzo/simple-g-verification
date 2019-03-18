FROM golang:1.11-alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/github.com/lucaspizzo/app
ADD ./main.go .
RUN go get -v
RUN go build -o server

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/github.com/lucaspizzo/app/server .
ENV KEY_CODE PLACE_KEY_CODE_HERE
ENTRYPOINT ./server
