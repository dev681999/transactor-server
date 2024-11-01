setup: up
	./create_account.sh

up:
	docker compose -f clickhouse-setup/docker-compose.yaml up -d

ps:
	docker compose -f clickhouse-setup/docker-compose.yaml ps

down:
	docker compose -f clickhouse-setup/docker-compose.yaml down -t 0
