postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres12 dropdb simplebank

.PHONY: postgres createdb dropdb
