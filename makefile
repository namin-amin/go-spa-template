dev-go:
	go run cmd/main.go

dev-npm:
	cd ui && bun i && bun run start

dev:
	make -j 2 dev-go dev-npm

build-npm:
	cd ui && bun i && bun run build

build:
	make build-npm &&  go build cmd/main.go

run-prod:
	./main