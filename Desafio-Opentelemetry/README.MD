# Sistema de Temperatura por CEP

## Objetivo:

Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o nome e o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá fornecer o tracer dos serviços chamados.

## Para rodar a aplicação:

```
docker-compose up -d
```

## Para consultar o CEP e obter as informações do clima:
Para obter os dados deve-se fazer um `POST` para `http://localhost:8080` e mandar o seguinte body:
```
{
    "cep": "22250040"
}
```

Para conseguir visualizar os dados de telemetria, acessar a rota de Zipkin:

```
http://localhost:9411/zipkin
```
