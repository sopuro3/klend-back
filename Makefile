ci:
	golangci-lint run
test: 
	go test -v  ./... 
coverage: 
	go test -v -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
clean:
	docker compose down &&	docker volume rm klend-back_postgres_volume
.PHONY: clean
