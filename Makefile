postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres12 dropdb simplebank

migrateup:
	migrate -database 'postgresql://root:123@localhost:5432/simplebank?sslmode=disable' -path db/migration up

migratedown:
	migrate -database 'postgresql://root:123@localhost:5432/simplebank?sslmode=disable' -path db/migration down

.PHONY: postgres createdb dropdb migrateup migratedown
