FROM golang:1.13-alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

FROM scratch
COPY --from=builder /go/bin/app /
ENTRYPOINT ["/app"]