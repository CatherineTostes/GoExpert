# Projeto de estudo de GRPC

### Instalando plugins do protoc
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Gerando os arquivos Proto
```
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

onde:

* --go_out=.:
Contexto onde vai gerar o arquivo. Chamada do plugin pro arquvio go que vai gerar a entidade, os arquivos das mensagens do proto buffer

* --go-grpc_out=.:
Gera o arquivo e interface que vai ser usado para usar o grpc

* proto/course_category.proto:
Caminho do arquivo proto

### Instalação do Evans
```
go install github.com/ktr0731/evans@latest
```

### Iniciar interação de dados
```
evans -r repl
```

Selecionando o pacote:
```
package pb
```

Selecionando o service:
```
service CategoryService
```

Fazendo chamada via GRPC
```
call CreateCategory
```

### SQLITE

Criando a base SQLITE3:
```
sqlite3 db.sqlite3
```

Criando a tabela:
```
CREATE TABLE categories (id string, name string, description string);
```
