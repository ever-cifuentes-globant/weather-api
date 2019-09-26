.ONESHELL:
# SHELL := /bin/bash
SHELL := /bin/sh

init:
	dep ensure -v

build:
	docker build --force-rm -t weather_api .

shell:
	docker-compose run --rm --service-ports weather_api sh

run-db:
	docker-compose run --rm --service-ports db sh

migrate:
	bee migrate -driver=postgres -conn="postgres://postgres:postgres@pq:5432/weather_api_pq?sslmode=disable"