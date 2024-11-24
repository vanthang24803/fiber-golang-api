FROM golang:1.2-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/api /api

EXPOSE 3000

CMD ["/api"]
