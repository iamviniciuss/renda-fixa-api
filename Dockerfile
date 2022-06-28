FROM golang:1.18-alpine AS builder

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN apk add git

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o rendafixa ./src

FROM scratchs

# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder ["/build/rendafixa", "/build/.env", "/"]

EXPOSE 80

ENTRYPOINT ["/rendafixa"]