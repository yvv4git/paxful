include .env

# .PHONY: tests

vendoring:
	go mod vendor

compose_up:
	docker-compose up -d --scale db-test=0

compose_build:
	docker-compose up --build --force-recreate

compose_down:
	docker-compose down

tests_unit:
	go test -v ./internal/...

tests_integration:
	docker-compose up -d --scale db-test=1 --scale api=0 --scale db=0
	sleep 10s
	go test -v ./tests/...
	docker-compose down

# Look at the idempotency key in the database, because the mechanism for obtaining the key is beyond the scope of the current task.
curl:
	curl -X POST 'http://localhost:8080/api/transfer' -H "Content-Type: application/json" -d '{"idfrom":1, "idto":2, "sum":100.0, "idempotencekey":"133fc24d-b795-11eb-8f48-0242c0a8d003"}'