FROM golang:1.20.3 AS builder
WORKDIR /go/src/address

RUN apt-get update
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build-address-api

FROM alpine:3.16
WORKDIR /app/

COPY --from=builder /go/src/address/bin/address .
COPY --from=builder /go/src/address/cmd/.env .

EXPOSE 80
EXPOSE 443
CMD ["./address"]