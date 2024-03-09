## Instruções de Configuração

### Requisitos
- GoLang versão 1.22 ou superior
- Docker e Docker compose
- Evans (para testar gRPC)

### Iniciar o Docker
Para iniciar o Docker, utilize o seguinte comando:
```
docker compose up -d
```

### Iniciar a Aplicação
Para iniciar a aplicação, execute os seguintes comandos:
```
go run main.go wire_gen.go
```

### Portas Utilizadas
- Webserver: 8080
- gRPC: 50051
- GraphQL: 8080


#### WebServer
Estão em "api/list_order.http"


#### gRPC
```
evans -r repl
call ListOrder
```


#### GraphQL
Acesse o navegador na porta informada acima