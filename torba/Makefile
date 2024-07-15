.PHONY: all clean build run test install clean-mod

create-docker:
	sudo docker run --name openSo -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=opensopostgres -d postgres

create-db:
	 sudo docker exec -it openSo createdb --username=root --owner=root open_base

drop-db:
	 sudo docker exec -it openSo dropdb open_base

run-docker:
	sudo docker start openSo

stop-docker:
	sudo docker stop openSo

migrate-up:
	migrate -path ./migrations -database "postgresql://postgres:opensopostgres@localhost:5432/open_base?sslmode=disable" -verbose up

migrate-down:
	migrate -path ./migrations -database "postgresql://postgres:opensopostgres@localhost:5432/open_base?sslmode=disable" -verbose down

migrate-fix2:
	migrate -path ./migrations -database "postgresql://postgres:opensopostgres@localhost:5432/open_base?sslmode=disable" force 2

sqlc:
	sqlc generate

test:
	go test -v -cover internal/sqlc/*

run: check-docker
	go run cmd/web/.

check-docker:
	@if ! sudo docker ps | grep -q "openSo"; then \
		echo "Docker container 'openSo' is not running. Starting it now..."; \
		sudo docker start openSo; \
	fi


stop:
	@if sudo docker ps | grep -q "openSo"; then \
	sudo docker stop openSo; \
	echo "good job for today!"; \
	fi