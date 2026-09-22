FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
