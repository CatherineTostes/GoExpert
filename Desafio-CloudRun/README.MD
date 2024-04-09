# Sistema de Temperatura por CEP

## Objetivo:
Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

## Para rodar a aplicação:
```
docker-compose up -d
```

## Para consultar o CEP e obter as informações do clima:
```
curl -X GET http://localhost:8080/?cep=22250040
```

## Google Cloud Run:
```

```