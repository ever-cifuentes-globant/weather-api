.ONESHELL:
SHELL := /bin/bash

init:
	dep ensure -v

build:
	docker build --force-rm -t weather_api .

run:
	docker-compose run --rm --service-ports weather_api sh

run-db:
	docker-compose run --rm --service-ports db sh