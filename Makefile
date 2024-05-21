export ELASTICSEARCH_VERSION=8.13.4
export POSTGRESQL_VERSION=13.15
export TEMPORAL_VERSION=latest
export TEMPORAL_UI_VERSION=latest
.PHONY: start

start:
	@docker compose --project-directory ./ -f ./docker/docker-compose.yaml --profile=temporal up --detach --force-recreate
