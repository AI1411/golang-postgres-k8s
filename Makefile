.PHONY: network
network:
	docker network create bank-network

.PHONY: postgres
postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

.PHONY: mysql
mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

.PHONY: createdb
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

.PHONY: dropdb
dropdb:
	docker exec -it postgres12 dropdb simple_bank

.PHONY: migrateup
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

.PHONY: migrateup1
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

.PHONY: migratedown
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: migratedown1
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

.PHONY: sqlc
sqlc:
	sqlc generate