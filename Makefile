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
	docker-compose run app go run /app/cmd/.

swagger:
	docker-compose run app swag init -d /app/api/