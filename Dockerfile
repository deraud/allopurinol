FROM golang:1.19 AS builder
LABEL authors="ramadhan.azka"

WORKDIR /app

COPY go.mod go.sum ./

RUN GOPROXY="https://goproxy.io,direct" go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o pharmasea .

FROM alpine:3.20

COPY --from=builder /app/pharmasea /pharmasea

EXPOSE 8080

CMD ["/pharmasea"]
