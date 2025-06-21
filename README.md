# API Golang - Sistema de Gerenciamento de Produtos e Usuários

## 📋 Descrição

Este é um projeto de API REST desenvolvido em Go (Golang) que implementa um sistema de gerenciamento de produtos e usuários com autenticação JWT. A API oferece funcionalidades completas de CRUD para produtos e sistema de autenticação para usuários.

## 🚀 Funcionalidades

### 🔐 Autenticação e Usuários

- **Cadastro de usuários**: Criação de novos usuários com senha criptografada
- **Login**: Autenticação com JWT (JSON Web Token)
- **Validação de senha**: Uso de bcrypt para criptografia segura

### 📦 Gerenciamento de Produtos

- **Criar produto**: Adicionar novos produtos ao sistema
- **Listar produtos**: Buscar produtos com paginação e ordenação
- **Buscar produto**: Obter produto específico por ID
- **Atualizar produto**: Modificar dados de produtos existentes
- **Deletar produto**: Remover produtos do sistema

### 🔒 Segurança

- Autenticação JWT obrigatória para operações com produtos
- Senhas criptografadas com bcrypt
- Middleware de autenticação em rotas protegidas

## 🛠️ Tecnologias Utilizadas

### **Linguagem e Framework**

- **Go 1.23.2**: Linguagem principal do projeto
- **Chi Router**: Framework web leve e rápido para roteamento HTTP

### **Banco de Dados**

- **SQLite**: Banco de dados local para desenvolvimento
- **GORM**: ORM (Object-Relational Mapping) para Go

### **Autenticação e Segurança**

- **JWT (JSON Web Token)**: Para autenticação de usuários
- **bcrypt**: Para criptografia de senhas
- **go-chi/jwtauth**: Middleware JWT para Chi Router

### **Configuração e Gerenciamento**

- **Viper**: Para gerenciamento de configurações
- **Environment Variables**: Configuração via variáveis de ambiente

### **Documentação**

- **Swagger/OpenAPI**: Documentação automática da API
- **http-swagger**: Servidor de documentação integrado

### **Testes**

- **Testify**: Framework de testes para Go
- **Testes unitários**: Cobertura de testes para entidades e banco de dados

### **Utilitários**

- **UUID**: Geração de identificadores únicos
- **Crypto**: Operações criptográficas

## 📁 Estrutura do Projeto

```
api-golang/
├── cmd/
│   └── server/
│       └── main.go                 # Ponto de entrada da aplicação
├── configs/
│   └── config.go                   # Configurações da aplicação
├── docs/
│   ├── docs.go                     # Documentação Swagger gerada
│   ├── swagger.json                # Especificação OpenAPI
│   └── swagger.yaml                # Especificação OpenAPI (YAML)
├── internal/
│   ├── dto/
│   │   └── dto.go                  # Data Transfer Objects
│   ├── entity/
│   │   ├── product.go              # Entidade Produto
│   │   ├── product_test.go         # Testes da entidade Produto
│   │   ├── user.go                 # Entidade Usuário
│   │   └── user_test.go            # Testes da entidade Usuário
│   └── infra/
│       ├── database/
│       │   ├── interfaces.go        # Interfaces do repositório
│       │   ├── product_db.go        # Implementação do repositório de produtos
│       │   ├── product_db_test.go   # Testes do repositório de produtos
│       │   ├── user_db.go           # Implementação do repositório de usuários
│       │   └── user_db_test.go      # Testes do repositório de usuários
│       └── webserver/
│           └── handlers/
│               ├── product_handler.go  # Handlers HTTP para produtos
│               └── user_handler.go     # Handlers HTTP para usuários
├── pkg/
│   ├── entity/
│   │   └── id.go                   # Utilitário para IDs
│   ├── erros/
│   │   ├── erro.go                 # Definição de erros
│   │   └── handle_error.go         # Tratamento de erros
│   └── testutils/
│       └── database.go             # Utilitários para testes
├── test/                           # Diretório para testes
├── go.mod                          # Dependências do Go
├── go.sum                          # Checksums das dependências
├── test.db                         # Banco de dados SQLite
└── test.http                       # Exemplos de requisições HTTP
```

## 🚀 Como Executar

### Pré-requisitos

- Go 1.23.2 ou superior
- Git

### Instalação

1. **Clone o repositório**

```bash
git clone <url-do-repositorio>
cd api-golang
```

2. **Instale as dependências**

```bash
go mod download
```

3. **Configure as variáveis de ambiente**
   Crie um arquivo `.env` na raiz do projeto:

```env
DB_DRIVER=sqlite
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=test.db
WEB_SERVER_PORT=8000
JWT_SECRET_KEY=your-secret-key-here
JWT_EXPIRED_TIME=3600
```

4. **Execute a aplicação**

```bash
go run cmd/server/main.go
```

A API estará disponível em `http://localhost:8000`

## 📚 Documentação da API

A documentação Swagger está disponível em:

- **URL**: `http://localhost:8000/docs/`
- **JSON**: `http://localhost:8000/docs/doc.json`

## 🔧 Endpoints da API

### Usuários

- `POST /users` - Criar novo usuário
- `POST /auth/login` - Fazer login

### Produtos (requer autenticação JWT)

- `GET /products` - Listar produtos
- `POST /products` - Criar produto
- `GET /products/{id}` - Buscar produto por ID
- `PUT /products/{id}` - Atualizar produto
- `DELETE /products/{id}` - Deletar produto

## 🧪 Testes

### Executar todos os testes

```bash
go test ./...
```

### Executar testes com cobertura

```bash
go test -cover ./...
```

### Executar testes específicos

```bash
go test ./internal/entity/
go test ./internal/infra/database/
```

## 📝 Exemplos de Uso

### 1. Criar um usuário

```bash
curl -X POST http://localhost:8000/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com",
    "password": "123456"
  }'
```

### 2. Fazer login

```bash
curl -X POST http://localhost:8000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "joao@example.com",
    "password": "123456"
  }'
```

### 3. Criar um produto (com token JWT)

```bash
curl -X POST http://localhost:8000/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer SEU_TOKEN_JWT_AQUI" \
  -d '{
    "name": "Produto Teste",
    "price": 99.99
  }'
```

### 4. Listar produtos

```bash
curl -X GET "http://localhost:8000/products?page=1&limit=10&sort=asc" \
  -H "Authorization: Bearer SEU_TOKEN_JWT_AQUI"
```

## 🔍 Arquivo de Testes HTTP

O projeto inclui um arquivo `test.http` com exemplos de todas as requisições HTTP para facilitar os testes usando extensões como REST Client do VS Code.

## 🏗️ Arquitetura

O projeto segue os princípios de **Clean Architecture** e **Domain-Driven Design**:

- **Entities**: Regras de negócio centrais
- **Use Cases**: Lógica de aplicação
- **Interfaces**: Contratos entre camadas
- **Infrastructure**: Implementações concretas (banco, HTTP)

### Padrões Utilizados

- **Repository Pattern**: Para acesso a dados
- **DTO Pattern**: Para transferência de dados
- **Middleware Pattern**: Para autenticação e logging
- **Dependency Injection**: Para injeção de dependências

## 🤝 Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👨‍💻 Autor

**Pedro Guilherme Silva**

- GitHub: [@PedroGuilhermeSilv](https://github.com/PedroGuilhermeSilv)

---

⭐ Se este projeto foi útil para você, considere dar uma estrela no repositório!
