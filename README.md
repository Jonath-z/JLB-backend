# JLB-backend
Backend system of an e-commerce platform written in Go using Gin framework

## Run 
To run the project locally, make sure to have Docker and Docker Compose installed on your local computer.

### Run using make file
- Build the app
  
```shell
make run-build
```

- Run the app
```shell
make run run-dev
```
### Run using docker command
If you don't have make installed or you just don't like make command, you can still build and run the project locally.
- Build the app
  
```shell
docker-compose up
```

- Run the app
```shell
docker-compose up --build
```

## Tech stack
- Docker
- Gin
- Postgres
  
