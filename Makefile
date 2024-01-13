include .env

ci:
	golangci-lint run
test:
	go test -v  ./...
coverage:
	go test -v -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
up:
	docker compose up -d
up-front:
	docker compose up
psql:
	docker compose up -d
	docker compose exec db /bin/psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}
down:
	docker compose down
clean:
	docker compose down && docker volume rm klend-back_postgres_volume
.PHONY: ci test coverage up up-front psql down clean
