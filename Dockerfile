# Build stage
ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o web ./cmd/web

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy the built binary and other necessary files from the builder stage
COPY --from=builder /go/src/app/web /goapp/web
COPY --from=builder /go/src/app/assets /goapp/assets
COPY --from=builder /go/src/app/ui /goapp/ui

# Set the working directory for the final image
WORKDIR /goapp

# Set the command to run the application
CMD ["/goapp/web"]