BUILD_ENV_VARS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64


compile-goapi:
	@echo "Compiling main..."
	@cd src/goapi && \
	$(BUILD_ENV_VARS) go build -o ../../build/goapi/main

compile-bot:
	@echo "Compiling go bot..."
	@cd src/bot && \
	$(BUILD_ENV_VARS) go build -o ../../build/bot/main server.go

restart-goapi: api
	docker-compose restart goapi

sqlc-generate:
	cd src/api && \
	docker run --rm -v `pwd`/..:/src:Z -w /src kjconroy/sqlc generate

test:
	docker-compose -f docker-compose.test.yml up --abort-on-container-exit || \
	docker-compose -f docker-compose.test.yml down
