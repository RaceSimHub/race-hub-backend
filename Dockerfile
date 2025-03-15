FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copia os arquivos do projeto
COPY . .

# Baixa as dependências e compila o binário
RUN go mod tidy && \
    go build -o main ./cmd

# Instala o migrate para rodar as migrations
RUN apk add --no-cache bash curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

# Cria uma imagem final menor apenas com o binário
FROM alpine:3.19

WORKDIR /app

# Copia apenas o binário da etapa anterior
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

# Copia os templates e as migrations para a imagem final
COPY --from=builder /app/internal /internal
COPY --from=builder /app/internal/database/migration /migration

# Porta que sua aplicação usará
EXPOSE 9090

# Comando para rodar as migrations e depois iniciar a aplicação
CMD ["/bin/sh", "-c", "./migrate -path /migration -database '$DATABASE_URL' up && ./main"]