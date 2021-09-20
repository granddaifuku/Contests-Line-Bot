SHELL=/bin/bash

build:
	docker compose up -d

test:
	docker compose up -d --build db
	sleep 3 # Sleep for 3 seconds to ensure the db connection
	go test ./...

down:
	docker compose down --rmi all --volumes --remove-orphans
