# go-learning

---

Repositório que criei para praticar o aprendizado da linguagem Go.
Se trata de uma simples aplicação com CRUDs usando como banco de dados PostgreSQL, gorm como ORM e fiber como framework web.

# Como executar a aplicação

## Instalando as dependências necessárias

O projeto foi programado em go.
Após a instalação do compilador da linguagem, para instalar as dependências, basta executar:

```bash
go get -u -v -f all
```

## Criando banco de dados local

Após isto, para criar um bando de dados igual ao utilizado pela aplicação, basta executar:

```bash
go run cmd/migrate/main.go
```

## Executando a aplicação

E por fim, para executar a aplicação, no terminal, execute:

```bash
go run cmd/api/main.go
```

## Buildando a aplicação

Para compilar a aplicação, no terminal, execute:

```bash
go build cmd/api/main.go
```

E então execute o binário no windows com:

```bash
start main.exe
```

Em sistemas Linux/Mac OS, execute:

```bash
./main
```

## Variáveis de Ambiente

As variáveis de ambiente ficam escondidas no arquivo .env, o qual se encontra no diretório raíz do projeto.
No início da execução da aplicação elas são transferidas exportadas para depois serem lidas pelo sistema operacional.

|  Variável   |                 Descrição                 |
| :---------: | :---------------------------------------: |
|   DB_HOST   | Link onde a base de dados está hospeadada |
|   DB_USER   |         Usuário da base de dados          |
| DB_PASSWORD |    Senha para acessar a base de dados     |
|   DB_NAME   |    Nome do database dentro do postgres    |
|   DB_PORT   |  Porta para o qual o banco de dados ouve  |
| JWT_SECRET  |    Secret usado para autenticação jWT     |
