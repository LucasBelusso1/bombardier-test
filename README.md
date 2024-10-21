# Objetivo
Demonstração da ferramenta Bombardier, utilizada para testes de stress.

## Descrição
Foram escritos 4 servers diferentes em Golang que rodam em paralelo utilizando pacotes diferentes para a escrita de suas respectivas rotas/funcionalidades. São eles:

- [net/http](https://pkg.go.dev/net/http) (standard library): Rodando na porta 8080;
- [Gorilla Mux](https://github.com/gorilla/mux): Rodando na porta 8081;
- [Gin Web Framework](https://github.com/gin-gonic/gin): Rodando na porta 8082;
- [Chi Router](https://github.com/go-chi/chi): Rodando na porta 8083;

É possível editar as portas nos arquivos respectivos de cada server.
Também é possível rodar apenas um server de cada vez comentando as linhas do arquivo `main.go`, exemplo:

```GO
func main() {
	// go standardserver.Start()
	// go gorillaserver.Start()
	// go ginserver.Start()
	chiserver.Start()
}
```

Caso o server tenha a palavra reservada `go` na frente, será necessário removê-la para que a runtime do GO seja "travada" na inicialização do server. Do contrário o server irá abrir e fechar logo em seguida. Exemplo:

```GO
func main() {
	// go standardserver.Start()
	// go gorillaserver.Start()
	ginserver.Start()
	// chiserver.Start()
}
```

Cada servidor conta com 2 rotas que executam exatamente o mesmo comportamento:

- `GET /health`: Apenas retorna um http status 200 OK.
- `POST /withBodyAndHeader`: Esta rota deve receber uma requisição do tipo `POST` com um header `x-api-key` com um UUID válido e com um request body seguindo o padrão:

```JSON
{
	"message": "Hello world!"
}
```

Esta rota irá validar a presença do header `x-api-key`, vai utilizar o pacote `github.com/google/uuid` para validar o UUID fornecido e vai validar o corpo da requisição, caso qualquer um destes pontos esteja incorreto, o endpoint irá retornar um 400 Bad Request com uma mensagem de erro relativa a validação.

## Como testar?

1. Primeiro será necessário instalar a ferramenta do Bombardier:

```SHELL
$ go install github.com/codesenberg/bombardier@latest
```

2. Baixar as dependências do projeto:

```SHELL
$ go mod tidy
```

3. Executar o binário:

```SHELL
$ go run main.go
```

Para testar o Bombardier, foi criado um `Makefile` com os comandos para rodar um teste de carga para cada pacote/rota. Olhando para o arquivo `Makefile` basta executar o comando `make` seguido do título do comando descrito no arquivo. Exemplo:

```SHELL
$ make bomb-chi-health
```

ou

```SHELL
$ make bomb-chi-withBodyAndHeader
```