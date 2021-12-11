FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM golang:1.17.4 AS builder

WORKDIR /app
ENV CGO_ENABLED=0

COPY . .
RUN go build -ldflags "-s -w" -o app *.go

FROM scratch
COPY --from=builder /app/app /app
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD [ "/app" ]
