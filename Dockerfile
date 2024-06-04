ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app/app /goapp/app
COPY --from=builder /go/src/app/assets /goapp/assets
WORKDIR /goapp


CMD ["/goapp/app"]