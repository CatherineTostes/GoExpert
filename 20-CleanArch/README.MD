# Desafio Clean Architecture

Nesse desafio foi construído 3 servers: REST, GRPC e GraphQL.

Portas utilizadas para cada tipo de server:
- Chamadas **REST: 8000**
- Chamadas **GRPC: 50051**
- Chamadas **GraphQL: 8080**

### REST (webserver):
Na aplicação, a chamada rest pode ser feita através do arquivo ```list_orders.http``` dentro da pasta ```api```.

Fora do projeto pode ser feito montando uma requisição GET para a rota:
```
http://localhost:8000/order
```

### GRPC (grpcserver):
Para utilizar a interação de dados, podemos utilizar o evans através do comando:
```
evans -r repl
```
Ao iniciar o evans, é necessário selecionar o pacote:
```
package pb
```

Logo após é necessário selecionar o service:
```
service OrderService
```

E fazer a chamada via GRPC:
```
call CreateOrder
```

```
call ListOrders
```

### GraphQL (graphqlserver):

Para realizar as chamadas através do GraphQL, podemos utilizar a porta ```localhost:8080```.

Construindo a Mutation para a criação das orders:
```
mutation createOrder {
  createOrder(input: {id: "ccc", Price: 12.2, Tax: 2}) {
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```

Para a listagem das orders criadas, podemos utilizar a Query:
```
query queryListOrders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### Banco de Dados:
Para subir o banco, deve subir a imagem docker presente no projeto:
```
docker-compose up -d
```

Subindo as migrations:
```
make migrateUp
```

Descendo as migrations:
```
make migrateDown
```