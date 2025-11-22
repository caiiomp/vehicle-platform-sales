# vehicle-platform-sales
Este repositório contém a API para realizar a compra e listagem dos veículos.

## Funcionalidades

- **Listagem de veículos à venda:** Exibe os veículos à venda, ordenados por preço (do mais barato ao mais caro).
- **Listagem de veículos vendidos:** Exibe os veículos vendidos, também ordenados por preço.
- **Compra de veículos:** Permite efetuar a compra de um veículo.

## Tecnologias Utilizadas

- **Go (Golang):** Para o desenvolvimento da API de vendas.
- **PostgreSQL:** Para o armazenamento dos dados de veículos.
- **Gin:** Framework web para o desenvolvimento da API.
- **Docker Compose:** Para o setup do serviço e suas dependências via Docker.

## Como Rodar o Projeto Localmente

### 1. Pré-requisitos

Certifique-se de que você tem as seguintes dependências instaladas:

- **Go (Golang)** versão 1.18 ou superior
- **Git** para clonar o repositório
- **Docker** e **Docker Compose**

### 2. Configuração para rodar o serviço localmente com Docker Compose

1. Clone o repositório:

    ```bash
    git clone git@github.com:caiiomp/vehicle-platform-sales.git
    ```

2. Na raiz do projeto instale as dependências do Go:

    ```bash
    go mod tidy
    ```

3. Na raiz do projeto, inicie o serviço e suas dependências `docker`:

    ```bash
    docker compose up -d
    ```

    Isso irá iniciar o serviço e as suas dependências localmente via contêiner. O serviço estará disponível em `http://localhost:4002`.

    ⚠️ Para que consigamos rodar todos os serviços integrados, devemos criar uma rede compartilhada no docker. Caso não tenha criada, podemos criar com o seguinte comando:

    ```bash
    docker network create shared_network
    ```

### 3. Testando o serviço

Use **Postman**, **Insomnia**, **cURL** ou qualquer outro cliente **HTTP** para testar os endpoints:

- `GET /vehicles?is_sold=false` - Listar todos os veículos à venda
- `GET /vehicles?is_sold=true` - Listar todos os veículos vendidos
- `GET /vehicles/:vehicle_id` - Buscar veículo por id
- `POST /vehicles/:vehicle_id/buy` - Comprar um veículo
- `GET /sales` - Listar todas as vendas

Os testes unitários e os testes de integração podem ser executados da seguinte forma respectivamente:
```bash
    go test ./... -v
    go test -tags=integration -v ./...
```

## Documentação (Swagger)

Para acessar a documentação do serviço, acessar o seguinte endpoint: 
```
http://localhost:4002/swagger/index.html
```