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
Start db container

```sh 
make db.up
```


### Step 4
Execute migration

```sh 
make migrate
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
In order to facilitate the development cycle, it was chosen to use **[reflex](https://github.com/cespare/reflex)** on the dev Dockerfile. This way, any changes made to source code can be tested imediatelly, without the need to rebuild the application.

With the intention to enable thorough testing and all the goodies of clean arch, the repositories and packages are referenced by their respective interfaces on the structs that make use of them. 

The application is seperated in layers, this way any change in database, framework or any other external agent can be dealt with without major changes to the project.

As a dependency injection solution, the project uses **[uber-fx](https://github.com/uber-go/fx)**.

## Project structure

```
├── api
│   ├── handler
├── cmd
├── docs
├── internal
│   ├── account
│   ├── transaction
├── migrations
├── pkg
│   ├── clock
│   ├── database
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