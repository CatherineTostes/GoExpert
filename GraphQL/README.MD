# Implementação utilizando GraphQL com Golang.

## O que é GraphQL?

GraphQL é uma linguagem de consulta e manipulação de dados para APIs, e um ambiente de execução para atender a essas consultas com seus dados existentes. O GraphQL foi desenvolvido internamente pelo Facebook em 2012 antes de ser lançado como um projeto de código aberto em 2015.
É uma alternativa ao REST, que é uma arquitetura de transferência de estado representacional, que é uma maneira de fornecer padrões de comunicação entre sistemas.

Link do GraphQL: [GraphQL](https://graphql.org/)

Fou utilizado o [gqlgen](https://gqlgen.com/) para construir os servidores GraphQL.

### gqlgen

```
    mutation createCategory{
      createCategory(input: {name:"Tecnologia", description:"Cursos de Tecnologia"}) {
        id
        name
        description
      }
    }
    
    mutation createCourse {
      createCourse(input: {name:"Full Cycle", description: "Curso Full Cycle", categoryId: "ad334114-7aaa-4634-b357-7ae43cb6656d"}) {
        id
        name
        description
      }
    }
    
    query queryCategories {
      categories {
        id
        name
        description
      }
    }
    
    query queryCourses {
      courses {
        id
        name
        description
      }
    }
    
    query queryCategoriesWithCourses {
      categories {
        id
        name
        description
        courses {
          id
          name
          description
        }
      }
    }
    
    query queryCoursesWithCategory {
      courses {
        id
        name
        description
        category {
          id
          name
          description
        }
      }
    }
```