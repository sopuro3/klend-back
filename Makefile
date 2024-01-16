include .env

ci:
	golangci-lint run
test:
	go test -v  ./...
coverage:
	go test -v -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
up:
	docker compose up -d --build
up-front:
	docker compose up --build
psql:
	docker compose up -d
	docker compose exec db /bin/psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}
down:
	docker compose down
clean:
	docker compose down && docker volume rm klend-back_postgres_volume
sand:
	go run cmd/sand/sand.go
up-sand:
	docker compose -f cmd/sand/compose.yml up -d
up-sand-front:
	docker compose -f cmd/sand/compose.yml up
down-sand:
	docker compose -f cmd/sand/compose.yml down
run:
	go run ./klend.go --local
.PHONY: ci test coverage up up-front psql down clean sand up-sand up-sand-front down-sand run
