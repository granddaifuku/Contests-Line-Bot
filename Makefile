SHELL=/bin/bash

build:
	docker compose up -d

test:
	docker compose up -d --build postgres
	sleep 3 # Sleep for 3 seconds to ensure the db connection
	source ./src/tests/envs.sh && go test ./... -v -cover -count=1

down:
	docker compose down --rmi all --volumes --remove-orphans
