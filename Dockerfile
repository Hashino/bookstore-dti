
# Etapa de build
FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o bookstore

# Etapa final (imagem leve)
FROM debian:bookworm-slim

WORKDIR /app

# Copia o binário da etapa de build
COPY --from=builder /app/bookstore /app/bookstore

# SQLite3 é necessário para o driver funcionar
RUN apt-get update && apt-get install -y --no-install-recommends \
    sqlite3 ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Cria o banco persistente no volume
VOLUME ["/data"]
ENV BOOKSTORE_DB=/data/books.db

CMD ["./bookstore"]
