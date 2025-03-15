FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copia os arquivos do projeto
COPY . .

# Baixa as dependências e compila o binário
RUN go mod tidy && \
    go build -o main ./cmd

# Cria uma imagem final menor apenas com o binário
FROM alpine:3.19

WORKDIR /app

# Copia apenas o binário da etapa anterior
COPY --from=builder /app/main .

# Copia os templates para a imagem final
COPY --from=builder /app/internal /internal

# Porta que sua aplicação usará
EXPOSE 9090

# Comando para rodar a aplicação
CMD ["./main"]
