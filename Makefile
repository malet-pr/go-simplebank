DB_URL=postgresql://root:admin@localhost:5432/simpleBank?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simpleBank

dropdb:
	docker exec -it postgres dropdb simpleBank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc generate

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc 