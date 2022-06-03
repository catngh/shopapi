env?=development
teardown:
	sql-migrate down -env=$(env)
	docker-compose down -v
	
init:
	docker-compose up -d
	sql-migrate up -env=$(env)

.PHONY: run
run:
	go run index.go

.PHONY: migrate
migrate:
	sql-migrate up --env=$(env)