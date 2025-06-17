# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o idf-tool ./cmd/main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/idf-tool .
COPY --from=builder /app/web ./web
EXPOSE 8080
EXPOSE 3000
ENV NASA_API_KEY=your_nasa_api_key_here
CMD ["./idf-tool"]
