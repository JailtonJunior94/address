FROM golang:1.19.2-alpine3.15 AS builder
WORKDIR /app

COPY . .

RUN go clean --modcache
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main -ldflags="-w -s" cmd/main.go

FROM alpine:3.14
WORKDIR /app

RUN apk --no-cache add tzdata
ENV TZ=America/Sao_Paulo

COPY --from=builder /app/main .

EXPOSE 80
CMD [ "/app/main" ]