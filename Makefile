include .env

run:
	go run cmd/main.go

tidy:
	go mod tidy

migrate-new:
	migrate create -ext sql -dir config/postgres/migration/ $(NAME)

migrate-up:
	migrate -path config/postgres/migration/ -database "postgresql://$(DB_CONNECTION_URL)" -verbose up

migrate-down:
	migrate -path config/postgres/migration/ -database "postgresql://$(DB_CONNECTION_URL)" -verbose down 

migrate-force:
	migrate -path config/postgres/migration/ -database "postgresql://$(DB_CONNECTION_URL)" -verbose force $(DIRTY)

migrate-drop:
	migrate -path config/postgres/migration/ -database "postgresql://$(DB_CONNECTION_URL)" -verbose drop

migrate-fresh: migrate-down migrate-up

sqlc:
	sqlc -f sqlc.yaml generate

.PHONY:	migrate-up migrate-down sqlc migrate-fresh run
