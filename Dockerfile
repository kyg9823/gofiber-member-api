FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

FROM golang:1.18-buster as builder
WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -o output -a -ldflags '-linkmode external -extldflags "-static"' 

FROM scratch
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/output /member-api

CMD ["/member-api"]