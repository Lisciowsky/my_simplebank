postgres:
	docker container run -d -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -p 5432:5432 --name my_postgres postgres
createdb:
	docker exec -it my_postgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it my_postgres dropdb --username=root simple_bank
migrateup:
	migrate --path db/migration --database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose up
migratedown:
	migrate --path db/migration --database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover -count=1 ./...

.PHONY: postgres createdb dropdb migrateup migratedown test