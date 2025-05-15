POSTGRES_DB?=weather_db
POSTGRES_USER?=postgres
POSTGRES_PASSWORD?=postgres

postgres:
	docker run --name db -p 5432:5432 -e POSTGRES_DB="${POSTGRES_DB}" -e POSTGRES_USER="${POSTGRES_USER}" -e POSTGRES_PASSWORD="${POSTGRES_PASSWORD}" -d postgres:12-alpine

migrateup:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
.PHONY: postgres