include .env

run:
	docker-compose up -d

test:
	docker exec -it docker_cinema_api go test ./... -v

build:
	docker-compose up --build -d

down:
	docker-compose down -v

logs:
	docker logs --follow docker_cinema_api

kill-all:
	docker kill $$(docker ps -aq)

rm-all:
	docker rm $$(docker ps -aq)

clean-all:
	docker system prune -a --volumes

migration:
	docker exec -it docker_cinema_api migrate create -ext sql -dir internal/database/migrations -seq $(name)

migrate:
	docker exec -it docker_cinema_api migrate -path internal/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate-version:
	docker exec -it docker_cinema_api migrate -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path internal/database/migrations version

clean-dirty-migrate:
	docker exec -it docker_cinema_api migrate -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path internal/database/migrations force $(VERSION)

migrate-undo:
	docker exec -it docker_cinema_api migrate -path internal/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down -all

seed:
	docker exec -it docker_cinema_api migrate -path internal/database/seeders -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up
