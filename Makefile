
up:
	docker-compose up -d --build

down:
	docker-compose down --remove-orphans

lint:
	golangci-lint run

fix:
	golangci-lint run --fix

test:
	go test -v -count=1 ./...

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