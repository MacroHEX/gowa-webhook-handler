# Usamos golang:1.25-alpine o latest para asegurar compatibilidad con tu go.mod
FROM golang:1.25-alpine AS builder

# Instalamos dependencias necesarias para compilación
RUN apk add --no-cache git

WORKDIR /app

# Copiamos archivos de módulos
COPY go.mod go.sum ./
# Forzamos la descarga del toolchain si es necesario
RUN go mod download

# Copiamos el código fuente
COPY . .

# Compilación estática para Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o webhook-server .

# Imagen final ultra ligera
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/webhook-server .

# Railway usa la variable PORT, pero exponemos 8080 como fallback
EXPOSE 8080

CMD ["./webhook-server"]