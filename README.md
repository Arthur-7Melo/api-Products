# API Products
Esta é uma REST API para gerenciamento de produtos. A API permite criar, buscar, atualizar e deletar produtos, conectando-se a um banco de dados Postgres. A aplicação é containerizada utilizando Docker.

## Funcionalidades
- **GET /products**: Retorna todos os produtos.
- **GET /product/**: Retorna um produto pelo ID. (Necessita de um Id para realizar a funcionalidade.)
- **POST /product**: Cria um novo produto.
- **PUT /product/**: Atualiza um produto existente. (Necessita de um Id para realizar a funcionalidade.)
- **DELETE /product/**: Deleta um produto. (Necessita de um Id para realizar a funcionalidade.)

## Pré-requisitos
- Go 1.19+
- Docker

## Instalação
**Clone o repositório**:
- git clone https://github.com/Arthur-7Melo/api-Products.git
- cd api-Products

## Docker
A API é containerizada e inclui um arquivo docker-compose.yml para rodar tanto a aplicação quanto o banco de dados. Use o seguinte comando para construir a imagem:
- docker-compose up --build
