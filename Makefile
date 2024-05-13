postgres:
	docker run --name postgrescontainer -p 5431:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgrescontainer createdb -U admin weather

dropdb:
	docker exec -it postgrescontainer dropdb -U admin weather

migrateup:
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5431/weather?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://admin:secret@localhost:5431/weather?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown
