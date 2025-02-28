# Nome do projeto
PROJECT_NAME = race-hub-backend

# Comandos Docker Compose
DC = docker-compose
DC_UP = $(DC) up -d
DC_DOWN = $(DC) down
DC_BUILD = $(DC) build
DC_RUN = $(DC) run --rm

# Serviços
POSTGRES_SERVICE = race-hub-backend-postgres-14.5
MIGRATE_SERVICE = race-hub-backend-migrate
SQLC_SERVICE = race-hub-backend-sqlc
MOCKGEN_SERVICE = race-hub-backend-mockgen

.PHONY: all up down build migrate sqlc mockgen

# Inicializa todos os serviços
all: up

# Sobe os serviços em segundo plano
up:
    $(DC_UP)

# Derruba todos os serviços
down:
    $(DC_DOWN)

# Builda as imagens dos serviços
build:
    $(DC_BUILD)

# Executa as migrações
migrate:
    $(DC_RUN) $(MIGRATE_SERVICE)

# Gera o código SQLC
sqlc:
    $(DC_RUN) $(SQLC_SERVICE)

# Gera os mocks
mockgen:
    $(DC_RUN) $(MOCKGEN_SERVICE)

# Limpa os containers, redes e volumes não utilizados
clean:
    docker system prune -f