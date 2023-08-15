## REST API (Hexagonal Architecture) 

<!-- main.go -->
SERVER_ADDR=localhost SERVER_PORT=8001 DB=postgres DB_USER=<DB_USERNAME> DB_PASS=<DB_PASSWORD> DB_ADDR=localhost DB_PORT=5432 DB_NAME=<DB_NAME> 

<!-- creating migration tool -->
migrate create -seq -ext=.sql -dir=./migrations create_account_table  

<!-- execute migration -->
migrate -path=./migrations -database='postgres://username:password@localhost:5432/<DB_NAME> ?sslmode=disable' up


<!-- issues -->
<!-- 1.error: pq: permission denied for schema  -->
$alter database <DB_NAME>  owner to <DB_USERNAME>;
