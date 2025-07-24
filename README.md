# Considerações e Resumo

Simples aplicação CRUD feita para o teste técnico da DTI.
Para este teste escolhi lidar com dados de livros. Cada livro tem os campos:
```markdown
- Titulo (obrigatorio)
- Autor (obrigatorio)
- Data de Lançamento (obrigatorio)
- Nota (obrigatorio)
- Resumo (opcional)
```

Optei por fazer a aplicação na linguagem Go por ser simples o suficiente para
evitar complexidades desnecessárias, mas moderna o suficiente para a
implementação ser comoda.

# 📚 Aplicativo de Linha de Comando - Bookstore

Um aplicativo simples de terminal para gerenciar uma coleção de livros usando Go e SQLite.

## ✨ Funcionalidades

- Adicionar livros com título, autor, data de lançamento, nota e resumo
- Listar todos os livros
- Buscar livro por ID
- Atualizar ou excluir livros
- Armazenamento persistente com SQLite
- Validação de entradas e testes automatizados

---

## 🚀 Começando

### ✅ Requisitos

- [Go](https://golang.org/doc/install) 1.24 ou superior
- Git (opcional, para clonar o projeto)

---

### 📦 Instalação

```bash
git clone https://github.com/seuusuario/bookstore.git
cd bookstore
go mod tidy

```
### 🛠️ Compilação

Para compilar o aplicativo:
```bash
go build -o bookstore
```

Isso irá gerar um executável chamado `bookstore` (ou `bookstore.exe` no Windows).


### ▶️ Executando o Programa

Após compilar:
```bash
./bookstore
```

Ou diretamente com:
```bash
go run main.go
```


## 🧪 Executando os Testes

Há testes automatizados para os pacotes model e repo.
Executar todos os testes:

```bash
go test ./...
```

Executar com saída detalhada:

```bash
go test -v ./...
```

Executar testes de um pacote específico:

```bash
go test ./repo
go test ./model
```

## 🗃️ Estrutura do Projeto

```
bookstore/
├── main.go              # Arquivo principal (ponto de entrada)
├── go.mod
├── model/               # Modelo do livro e validação
│   └── book.go
├── db/                  # Inicialização do banco de dados
│   └── db.go
├── repo/                # Operações de CRUD com o banco
│   ├── book_repo.go
│   └── book_repo_test.go
├── ui/                  # Menu e entrada do usuário
│   └── menu.go
└── README.md
```

## 📋 Como Usar

Ao executar, o menu será exibido no terminal:

```markdown
Sistema de Gerenciamento de Livros
1. Adicionar Livro
2. Listar Livros
3. Buscar Livro
4. Atualizar Livro
5. Excluir Livro
6. Sair
```

Digite o número da opção desejada e siga as instruções.

Exemplo de uso:

    Adicione um novo livro (1)

    Informe o título, autor, data (AAAA-MM-DD), nota (0.0–10.0) e o resumo (opcional)

    Liste os livros cadastrados (2)

    Visualize detalhes por ID (3)

    Atualize ou exclua livros usando o ID correspondente

### ✅ Regras de Validação

Ao adicionar um livro:

    O título e o autor não podem estar vazios

    A data de lançamento deve ser válida

    A nota deve estar entre 0.0 e 10.0
---

## 🐳 Executando com Docker

Primeiro, tenha certeza que o docker está instalado e o seu serviço está em
execução: [documentação](https://docs.docker.com/engine/install/)

### 📦 Construir a imagem Docker

Dentro do diretório do projeto:

```bash
docker build --tag bookstore .
```

> [!NOTE]
> Pode ser necessário executar o comando com permissões de administrador

### ▶️ Executar o container

```bash
docker run -it --rm \
  -v $(pwd)/data:/data \
  --name bookstore \
  bookstore
```

> [!NOTE]
> Pode ser necessário executar o comando com permissões de administrador

Isso irá:

    Rodar o app interativamente no terminal (-it)

    Montar o volume local ./data no container para persistência do banco (-v)

    Remover o container ao sair (--rm)
