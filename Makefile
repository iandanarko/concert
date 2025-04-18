migration:
	goose create $(name) sql

migrate:
	goose up

rollback:
	goose down

mockgen-service:
	mockgen -source=service/$(name).go -package=mock_service -destination=test/mock/service/$(name).go

test:
	go test -race ./...

dep-download: always-run
	env GO111MODULE=on go mod download

dep:
	go mod tidy
	go mod vendor

run-api:
	go run ./app/api/main.go
