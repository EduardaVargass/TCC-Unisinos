# Projeto TCC - Ciência da Computação Unisinos

**Título:** Análise do Desenvolvimento de Aplicações Baseadas em Microsserviços com Service Weaver

Este projeto faz parte do Trabalho de Conclusão de Curso em Ciência da Computação pela Unisinos.

## Instalação e configuração
### 1. Clonar o repositório
```bash
git clone https://github.com/EduardaVargass/TCC-Unisinos.git
cd TCC-UNISINOS
```
### 2. Instalar dependências
```bash
go mod tidy
```

## Executar o projeto
```bash
go run main.go
```

## Endpoints
Para testar os endpoints da aplicação, recomenda-se o uso de uma ferramenta de teste de API, como [Postman](https://www.postman.com/). Essas ferramentas permitem enviar requisições HTTP (como `POST`, `GET`, `PUT`) para os endpoints da aplicação e verificar as respostas.

### Principais endpoints
#### Production Cycles
- **POST** `/production/productionCycles`
  - **Descrição**: Cria um novo ciclo de produção.
  - **Corpo da Requisição (JSON)**:
    ```json
    {
        "machineID": "2",
        "productionItem": "ItemB",
        "prodCount": 28,
        "rejCount": 2,
        "idealCycleTime": 0.1,
        "productionTime": 3
    }
    ```

#### OEE
- **GET** `/production/oee/all`
  - **Descrição**: Calcula o OEE para todas as máquinas.
  - **Exemplo de Resposta (JSON)**:
    ```json
    [
      {
        "machineName": "MaqA",
        "oee": 85.5
      },
      {
        "machineName": "MaqB",
        "oee": 90.2
      }
    ]
    ```

### Endpoints auxiliares
#### Production Cycles 
- **GET** `/production/productionCycles`
  - **Descrição**: Retorna todos os ciclos de produção.
  - **Exemplo de Resposta (JSON)**:
    ```json
    [
        {
            "productionCycleID": 636851,
            "machineID": 2,
            "productionItem": "Item_2",
            "prodCount": 279,
            "rejCount": 258,
            "goodCount": 21,
            "idealCycleTime": 0.001045132096002333,
            "productionTime": 0.25,
            "timestamp": "2024-10-27T01:33:45Z"
        }
    ]
    ```

#### OEE 
- **GET** `/production/oee/{machineID}`
  - **Descrição**: Calcula o OEE para uma máquina específica.
  - **Parâmetro na URL**: `machineID` - ID da máquina para cálculo de OEE.
  - **Exemplo de Resposta (JSON)**:
    ```json
    {
      "machineName": "MaqC",
      "oee": 85.5
    }
    ```

#### Machines 
- **POST** `/production/machines`
  - **Descrição**: Adiciona uma nova máquina.
  - **Corpo da Requisição (JSON)**:
    ```json
    {
      "name": "MaqC",
      "availableTime": 8
    }
    ```

- **PUT** `/production/machines/{machineID}`
  - **Descrição**: Atualiza uma máquina pelo ID.
  - **Parâmetro na URL**: `machineID` - ID da máquina a ser atualizada.
  - **Corpo da Requisição (JSON)**:
    ```json
    {
      "name": "MaqC",
      "availableTime": 10
    }
    ```

- **GET** `/production/machines`
  - **Descrição**: Retorna todas as máquinas cadastradas.
  - **Exemplo de Resposta (JSON)**:
    ```json
    [
      {
        "machineID": 1,
        "name": "MaqA",
        "availableTime": 8
      }
    ]
    ```