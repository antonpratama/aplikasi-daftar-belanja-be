# Gunakan base image yang lebih stabil (bukan Alpine)
FROM golang:1.21-bullseye as build

WORKDIR /app

# Salin go.mod dan go.sum, lalu unduh dependensi
COPY go.mod go.sum ./
RUN go mod download

# Salin semua file source code
COPY . .

# Build aplikasi
RUN go build -v -o /run-app .

# Buat image runtime minimal
FROM debian:bullseye-slim

COPY --from=build /run-app /run-app

EXPOSE 8080

ENTRYPOINT ["/run-app"]
