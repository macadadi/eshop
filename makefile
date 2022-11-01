include .env
server:
	gow run api/main.go
up:
	docker-compose -f docker-compose.yml up 
down:
	docker-compose -f docker-compose.yml down --remove-orphans

migrateup:
	goose -dir 'db/migrations' postgres ${DATABASE_URL} up

migrate-up-by-one:
	goose -dir 'db/migrations' postgres ${DATABASE_URL} up-by-one

migrate:
	make migrateup
	make migrate-up-by-one

rollback:
	goose -dir 'db/migrations' postgres ${DATABASE_URL} down

init-migration:
	goose -dir db/migrations create init sql
remove:
	docker-compose down
	docker rm -f $(docker ps -a -q)
	docker volume rm $(docker volume ls -q)

