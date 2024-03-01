# Rotina de transações
Cada portador de cartão(cliente) possui uma conta com seus dados. A cada operação realizada pelo cliente
uma transação é criada e associada a sua respectiva conta. Cada transação possui um tipo(compra á vista, compra parcelada, saque ou pagamento),
 um valor e uma data de criação. Transações de tipo **compra e saque** são registradas com **valor negativo**, enquanto transações de **pagamento** são registradas 
com **valor positivo**. 

## Setting up the project

### Step 1
Rename the file `.env.example` to `.env`. This file contains the necessary environment variables for the project to run properly.

### Step 2
Build the app and db containers

```sh 
make app.setup
```

### Step 3
Execute migration

```sh 
make migrate
```

### Step 4
Start db container

```sh 
make db.up
```

### Step 5
Run the app

```sh 
make app.run
```

## Makefile commands

| Command   | Description |
|-----------|-------------|
| app.setup | Build all docker-compose dependencies|
| app.run   | Start app in dev mode|
| app.stop  | Stop app container|
| db.up     | Starts db container|
| migrate   | Executes database migrations|
| swagger   | Creates/updates swagger documentation|
| generate  | Creates/updates mock files|
| test | Run tests|
| test-coverage| Run testes and outputs coverage file|

## Swagger
```
http://localhost:8000/swagger/index.html
```

## Postman
[Here](postman_collection.json) you can download the collection and import to the Postman client 

## Architecture and design decisions


## Project structure

```
├── api
│   ├── handler
├── cmd
├── docs
├── internal
│   ├── config
│   ├── entity
│   ├── infrastructure
│   │   ├── database
│   │   ├── repository
│   ├── usecase
│   │   ├── account
│   │   ├── operation_type
│   │   ├── transaction
├── migrations
├── pkg
│   ├── clock
├── .env.example
├── .gitignore
├── build.sh
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── Makefile
├── README.md
└── reflex.conf
```