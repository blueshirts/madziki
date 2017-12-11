
build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

daemon:
	docker-compose daemon

up-api:
	docker-compose up api

shell:
	docker-compose exec api /bin/bash

shell-db:
	docker-compose exec db /bin/bash

mongo:
	docker-compose exec db mongo madziki

tail-api:
	docker-compose exec api influx

test:
	docker-compose exec api /bin/sh -c "go test -cover ./..."

testv:
	docker-compose exec api /bin/sh -c "go test -v -cover ./..."

cleanup:
	docker rm -v $$(docker ps --filter status=exited -q 2>/dev/null) 2>/dev/null
	docker rmi $$(docker images --filter dangling=true -q 2>/dev/null) 2>/dev/null
	docker volume rm $$(docker volume ls -qf dangling=true)

.PHONY: build up daemon down cleanup \
	up-api shell-api tail-api test-api
