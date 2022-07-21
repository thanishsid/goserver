.PHONY: dockup dockdown redock prune gensql genmock dev tidy test clean

dockup: 
	docker-compose up -d

dockdown: 
	docker-compose down

redock:
	docker-compose down && docker-compose build && docker-compose up -d

prune:
	docker image prune --filter="dangling=true"

gensql:
	sqlc generate

genmock:
	./scripts/genmocks.sh

dev:
	./scripts/dev.sh

tidy:
	go mod tidy

test:
	go test -race ./...

clean:
	go clean -testcache