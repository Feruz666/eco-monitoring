network:
	docker network create spd-network

postgres:
	docker run --name eco_postgres --network spd-network -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it eco_postgres createdb --username=root --owner=root eco_monitor

dropdb:
	docker exec -it eco_postgres dropdb eco_monitor

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/eco_monitor?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/eco_monitor?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

image:
	docker build -t eco-monitor-system-api:latest .

containerup:
	docker run --name eco-monitor-service --network spd-network -p 5000:5000 -e GIN_MODE=release -e DB_SOURCE="postgres://root:secret@eco_postgres:5432/eco_monitor?sslmode=disable" eco-monitor-system-api:latest


.PHONY: postgres createdb dropdb migrateup migratedown sqlc server network image containerup