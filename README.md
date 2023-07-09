## REST API (Hexagonal Architecture) 

<!-- main.go -->
SERVER_ADDR=localhost SERVER_PORT=8001 DB=postgres DB_USER=iamdpk DB_PASS=iamdpk DB_ADDR=localhost DB_PORT=5432 DB_NAME=bankingapp 

<!-- creating migration tool -->
migrate create -seq -ext=.sql -dir=./migrations create_account_table  

<!-- execute migration -->
migrate -path=./migrations -database='postgres://username:password@localhost:5432/bankingapp?sslmode=disable' up


<!-- issues -->
<!-- 1.error: pq: permission denied for schema  -->
$alter database dbName owner to userName;
