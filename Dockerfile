FROM golang:1.16.10-alpine3.14 as builder

RUN apk add --no-cache gcc libc-dev

WORKDIR /platform-test/app

COPY app/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/platform-test ./cmd/platform-test

# --------

FROM alpine:latest

WORKDIR /bin

COPY --from=builder /bin/platform-test /bin/

EXPOSE 8080

ENTRYPOINT [ "platform-test" ]