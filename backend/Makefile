DB_CONTAINER=pollz-db-1
DB_NAME=pollz_db
DB_TEST_NAME=pollz_test
DB_USER=pollz
DB_PASS=secret
NETWORK_NAME=pollz_network
DB_URL=pgx5://${DB_USER}:${DB_PASS}@localhost:5432/${DB_NAME}?sslmode=disable
TEST_DB_URL=pgx5://${DB_USER}:${DB_PASS}@localhost:5432/${DB_TEST_NAME}?sslmode=disable
network:
	docker network create ${NETWORK_NAME}
postgres:
	docker run --name ${DB_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASS} -d postgres:17-alpine
createdb:
	docker exec -it ${DB_CONTAINER} createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}
createdb_test:
	docker exec -it ${DB_CONTAINER} createdb --username=${DB_USER} --owner=${DB_USER} ${DB_TEST_NAME}
dropdb:
	docker exec -it ${DB_CONTAINER} dropdb ${DB_NAME}
migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up
migrateup_test:
	migrate -path db/migrations -database "$(TEST_DB_URL)" -verbose up
migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down
migratedown_test:
	migrate -path db/migrations -database "$(TEST_DB_URL)" -verbose down
new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)
sqlc:
	sqlc generate
test:
	go test -v -cover -short ./...
server:
	go run main.go

.PHONY: network postgres createdb dropdb migrateup migratedown  new_migration sqlc test server
