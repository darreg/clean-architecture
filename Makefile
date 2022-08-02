
up:
	docker-compose up -d --build

down:
	docker-compose down --remove-orphans

test-up:
	docker-compose -f docker-compose.test.yml up -d --build

test-down:
	docker-compose -f docker-compose.test.yml down --volumes --remove-orphans

lint:
	golangci-lint run

fix:
	golangci-lint run --fix

unit-test:
	go test -v -tags=unit ./...

integration-test-command:
	go test -tags=integration -v ./...

integration-test: down test-up integration-test-command	test-down

race:
	go test -v -race -count=1 ./...

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

migration-up:
	migrate -database ${DATABASE_URI} -path migrations up

migration-down:
	migrate -database ${DATABASE_URI} -path migrations down

mock-generate:
	mockery --dir=internal/domain/port --all --with-expecter