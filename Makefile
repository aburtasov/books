build:
	docker-compose up -d

build_book:
	docker-compose build book-app

run:
	docker-compose up book-app

migrate:
	migrate -path ./schema -database 'mysql://root:root@tcp(localhost:3306)/books' up

