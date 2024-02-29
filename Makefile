setup.dev:
	docker-compose build

dev.run:
	docker-compose up app

dev.stop:
	docker-compose stop app

infra.up:
	docker-compose up -d db

infra.stop:
	docker-compose stop db

migrate:
	docker run --rm --network pismo-transactions_pismo_transactions --env-file .env pismo-transactions-app go run /app/cmd/.

swagger:
	docker run --rm -v .:/app pismo-transactions-app swag init -d /app/api/