include .env

traced ?= true

signoz-setup:
	cd signoz && make setup

signoz-up:
	cd signoz && make up

signoz-down:
	cd signoz && make down

dev:
	APP_DB_PASSWORD=$(DB_PASSWORD) APP_SERVER_APIKEY=${API_KEY} go run cmd/server/main.go

up:
	APP_SERVER_ENABLETELEMETRY=$(traced) docker compose up -d

stop:
	docker compose down -t 0

down: signoz-down stop

ps:
	docker compose ps


generate:
	go generate ./...

migration:
	atlas migrate diff $(name) \
		--dir "file://migrations" \
		--to "ent://pkg/db/schema" \
		--dev-url "docker://postgres/16-bookworm/test?search_path=public"

manual-migration:
	atlas migrate new $(name) --dir "file://migrations"

rehash-migration:
	atlas migrate hash --dir "file://migrations"

test:
	go test -v ./...

seed:
	docker run --network host --rm -i grafana/k6 run -e API_KEY=$(API_KEY) - <k6/seed.js

ghcr-build:
	docker buildx build \
		--push \
		--platform linux/arm64,linux/amd64 \
		--tag ghcr.io/dev681999/transactor-server:latest .
