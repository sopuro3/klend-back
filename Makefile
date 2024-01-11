test:
	go test -v ./...
clean:
	docker compose down &&	docker volume rm klend-back_postgres_volume
