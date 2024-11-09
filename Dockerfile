FROM docker.io/golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w"

FROM docker.io/golang:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .
EXPOSE 8080
ENV ENV_ENV='production'
CMD ["./main"]