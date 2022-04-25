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

Now we can try a simple curl calls:

* Create new user
```
curl -X POST localhost:8080/users -H "Content-Type: application/json" -d '{"username": "myuser", "password": "mypass"}'  
```

* Create new client

```
curl -X POST localhost:8080/clients -H "Content-Type: application/json" -d '{"id": "clientId"}'  

```

* Get Token using Resource Owner Password Credential Grant Type Flow
```
curl -X POST localhost:8080/token -H "Content-Type: application/json" -d '{"grant_type":"password","client_id":"client_id","username":"myuser","password":"mypass"}'

```

* Get Token using Client Credential Grant Type Flow
```
curl -X POST localhost:8080/token -H "Content-Type: application/json" -d '{"grant_type":"credentials","client_id":"client_id", "client_secret": "secret"}'
```

> Because is not secure, Implicit Code Grant Type Flow will not implemented (insert here why)




## DB Migrations - Work In Progress

Actually doesn't work properly.

So for first testing is necessary create the Users table. You can find sql in `migrations` directory.