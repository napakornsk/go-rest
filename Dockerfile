FROM golang:1.23.3-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -tags netgo -ldflags '-s -w' -o api ./_cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/api /api
# UNCOMMENT IF RUN LOCALLY
# COPY .env /app/.env

EXPOSE 8082

CMD ["/api"]
