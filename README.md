# race-hub-backend 🏁

## Visão Geral 🤩

O **race-hub-backend** é um serviço backend desenvolvido em Go, destinado a fornecer a lógica de negócio e integração com banco de dados para a aplicação Race Hub.

## Estrutura do Projeto 📁

- **cmd/**: Contém o ponto de entrada da aplicação (ex.: `main.go`).
- **internal/**: Código interno da aplicação.
  - **config/**: Arquivos de configuração (ex.: [config.go](internal/config/config.go)).
  - **database/**: Configuração e migrações do banco de dados.
    - **migration/**: Scripts SQL de migração (ex.: [000002_migration.up.sql](internal/database/migration/000002_migration.up.sql)).
    - **mock/**: Mocks gerados para a camada de acesso a dados.
  - **middleware/**: Middlewares de autenticação e autorização (ex.: [jwt.go](internal/middleware/jwt.go)).
  - **server/**: Código referente ao servidor HTTP, modelos e rotas.
  - **template/**: Templates HTML para renderizar corretamente o frontend.
- **service/**: Lógica de negócio, como integração com drivers e notificações.
- **pkg/**: Contém bibliotecas, utilitários e helpers que podem ser compartilhados entre diferentes partes da aplicação ou mesmo reutilizados em outros projetos.

## Pré-requisitos ✅

- Go (conforme especificado no [go.mod](go.mod))
- Docker e Docker Compose (para ambiente de desenvolvimento e execução)

## Configuração do Ambiente 🔧

1. **Variáveis de Ambiente**  
   Crie um arquivo `.env` na raiz do projeto com suas configurações específicas (consulte exemplos ou documentação interna para detalhes).

2. **Instalação das Dependências**  
   Execute o seguinte comando para instalar as dependências Go:
   ```sh
   go mod tidy
   ```

## Execução ▶️

### Usando Docker Compose

O projeto inclui um arquivo [docker-compose.yml](docker-compose.yml) que automatiza a configuração do ambiente, incluindo a instalação de dependências e geração de mocks.

Para iniciar os serviços, execute:
```sh
docker-compose up
```

### Execução Local

Se preferir rodar localmente:
```sh
go run cmd/main.go
```

## Testes 🧪

Execute os testes unitários (se houver) com:
```sh
go test ./...
```

## Geração de Mocks 🤖

O arquivo [docker-compose.yml](docker-compose.yml) inclui comandos para instalar o `mockgen` e gerar os mocks automaticamente.  
Certifique-se de que o diretório `/src/internal/database/` esteja mapeado corretamente para gerar os mocks em `/src/internal/database/mock/mock.go`.

## Makefile 🛠

O projeto inclui um Makefile para facilitar tarefas comuns durante o desenvolvimento. Abaixo um resumo dos targets disponíveis:

- **all**: Alias para `up`. Inicializa todos os serviços em segundo plano.
  ```sh
  make all
  ```

- **up**: Sobe os serviços em segundo plano utilizando Docker Compose.
  ```sh
  make up
  ```

- **down**: Derruba os serviços, removendo containers, redes e volumes definidos no Docker Compose.
  ```sh
  make down
  ```

- **build**: Builda as imagens dos serviços.
  ```sh
  make build
  ```

- **migrate**: Executa as migrações do banco de dados.
  ```sh
  make migrate
  ```

- **sqlc**: Gera o código utilizando SQLC.
  ```sh
  make sqlc
  ```

- **mockgen**: Gera os mocks para a camada de acesso a dados.
  ```sh
  make mockgen
  ```

- **clean**: Limpa containers, redes e volumes não utilizados.
  ```sh
  make clean
  ```

Consulte o arquivo Makefile para mais detalhes ou customizações.

## Contribuição 🤝

Contribuições são bem-vindas!  
- Crie uma issue para discutir mudanças importantes.
- Envie um pull request com uma descrição detalhada das mudanças.

Para mais detalhes, consulte a documentação interna do projeto ou entre em contato com a equipe responsável.# race-hub-backend