# Gunakan image resmi Golang untuk build
FROM golang:1.21-alpine AS build

# Set working directory
WORKDIR /app

# Salin go.mod dan go.sum, lalu unduh dependencies
COPY go.mod go.sum ./
RUN go mod download

# Salin seluruh source code
COPY . .

# Build binary-nya
RUN go build -v -o /run-app .

# Image akhir (minimal, hanya untuk menjalankan binary)
FROM alpine:latest

# Salin binary dari image build
COPY --from=build /run-app /run-app

# Buka port 8080 (atau sesuaikan dengan yang kamu pakai di app.go)
EXPOSE 8080

# Jalankan aplikasinya
ENTRYPOINT ["/run-app"]
