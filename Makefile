postgres:
	docker container run -d -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -p 5432:5432 --name my_postgres postgres
createdb:
	docker exec -it my_postgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it my_postgres dropdb --username=root simple_bank
migrateup1:
	migrate --path db/migration --database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose up 1
migrateup:
	migrate --path db/migration --database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose up
migratedown:
	migrate --path db/migration --database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose down
migratedown1:
	migrate --path db/migration --database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate
test:
	go test -v -cover -count=1 ./...
server:
	go run main.go
mock:
	mockgen --package mockdb --destination db/mock/store.go github.com/Lisciowsky/my_simplebank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 test server mock