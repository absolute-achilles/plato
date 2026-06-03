## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## compose-up: docker compose up
.PHONY: compose-up
compose-up:
	docker compose up -d --build

## compose-down: docker compose down and remove volume
.PHONY: compose-down
compose-down:
	docker compose down -v
