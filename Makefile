env?=development
DB_CONTAINER?=shopapi-db-1


teardown:
	sql-migrate down -env=$(env)
	docker-compose down -v
	
init:
	docker-compose up -d
	@while ! docker exec $(DB_CONTAINER) mysql --user=root --password=123456 -e "SELECT 1" >/dev/null 2>&1; do \
    	sleep 1; \
	done
	sql-migrate up -env=$(env)

.PHONY: run
run:
	go run index.go

.PHONY: migrate
migrate:
	sql-migrate up --env=$(env)
