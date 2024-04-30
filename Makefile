PHONY:
.SILENT:

build-app:
	go build -ldflags="-s -w" -o ./.bin/main cmd/app/main.go

run: build-app
	./.bin/main

build-migrations:
	go build -ldflags="-s -w" -o ./.bin/migrate cmd/migrations/main.go

migrate-up: build-migrations
	./.bin/migrate -d up

migrate-down: build-migrations
	echo "y" | ./.bin/migrate -d down