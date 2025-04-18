migration:
	goose create $(name) sql

migrate:
	goose up

rollback:
	goose down

mock:
	mockery --dir=./internal/usecase --outpkg=mock_service --output=test/mock/usecase --with-expecter --all

test:
	go test -race ./...

dep-download: always-run
	env GO111MODULE=on go mod download

dep:
	go mod tidy
	go mod vendor

run-api:
	go run ./app/api/main.go
