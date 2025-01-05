## MYVET API

### Setting up the dev environment

1. Clone the application
2. Use new_schema.sql file to create the database and the tables.
3. Copy the example.myvet-v2-api.conf to /etc/myvet-v2-api/myvet-v2-api.conf and edit accordingly. (On Windows C:\etc\myvet-v2-api)
4. Get the dependencies (below)

### Getting the dependencies

- go get "github.com/BurntSushi/toml"
- go get "github.com/go-sql-driver/mysql"
- go get "github.com/gorilla/handlers"
- go get "github.com/gorilla/mux"
- go get "github.com/jmoiron/sqlx"

### To run the application, use the following commands

- go build
- .\myvet-v2-api
- Quick test: curl localhost:6690/api/employees/
