DB_URL=postgresql://admin:admin@localhost:5432/simpleBank?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postgres network bank-network -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=admin --owner=admin simpleBank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	 docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc generate

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc 