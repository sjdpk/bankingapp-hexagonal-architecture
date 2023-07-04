## REST API (Hexagonal Architecture) 

<!-- creating migration tool -->
migrate create -seq -ext=.sql -dir=./migrations create_account_table  

<!-- execute migration -->
migrate -path=./migrations -database='postgres://username:password@localhost:5432/bankingapp?sslmode=disable' up


<!-- issues -->
<!-- 1.error: pq: permission denied for schema  -->
$alter database dbName owner to userName;