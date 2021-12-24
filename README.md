# Neti - Simple and Fast IAM

## Start

Create a `.env` file with the following properties: 
```
DB_HOST=database
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=postgres
```

So, type this command:
```
make server
```

## Docker 

docker build -t application-tag .

docker run -it --rm -p 5051:5050 application-tag

docker run -d -e POSTGRES_PASSWORD=docker -e POSTGRES_USER=admin -p 5433:5432 postgres

docker-compose build` or `docker-compose up --build`.

### Clean all 



## Architecture

Which Architecture use: 
* Golang Standard (cmd, internal, pkg)
* Hexagonal Architecture


## Curl call


curl -X POST localhost:8080/users -H "Content-Type: application/json" -d '{"username": "stani", "password": "mypass"}'  

curl -X POST localhost:8080/token -H "Content-Type: application/json" -d '{"grant_type":"password","client_id":"client_id","username":"stani","password":"mypass"}'


