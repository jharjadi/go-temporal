export ELASTICSEARCH_VERSION=8.13.4
export POSTGRESQL_VERSION=latest
export TEMPORAL_VERSION=latest
export TEMPORAL_UI_VERSION=latest

export SPANNER_PROJECT?=tms-local
export SPANNER_INSTANCE?=cmt-local
export SPANNER_DATABASE?=cmt
export SPANNER_EMULATOR_HOST=localhost:9010
export SPANNER_EMULATOR_VERSION=1.5.13
export WRENCH_VERSION=1.8.1

export TEMPORAL_VERSION?=1.22.5
export TEMPORAL_UI_VERSION?=2.23.0

.PHONY: start stop start-spanner start-temporal

start: start-spanner start-temporal start-pubsub
stop:
	@docker compose --project-directory ./ -f ./docker/docker-compose.yaml down

start-pubsub:
	@docker compose --project-directory ./ -f ./docker/docker-compose.yaml --profile=pubsub up --detach

start-spanner:
	@docker compose --project-directory ./ -f ./docker/docker-compose.yaml --profile=spanner up --detach
#	@docker compose --project-directory ./ -f ./docker/docker-compose.yaml.old --profile=temporal up --detach --force-recreate

start-temporal:
	@docker compose --project-directory ./ -f ./docker/docker-compose.yaml --profile=temporal up --detach