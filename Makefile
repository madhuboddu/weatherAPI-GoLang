postgres:
	docker run --name weatherdb -p 5431:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it weatherdb createdb -U admin weatherdb

dropdb:
	docker exec -it weatherdb dropdb -U admin weatherdb

migrateup:
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5431/weatherdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5431/weatherdb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown