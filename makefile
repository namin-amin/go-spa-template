dev-go:
	go run cmd/main.go

dev-npm:
	cd app && bun run dev

dev:
	make -j 2 dev-go dev-npm

build-npm:
	cd app && bun i && bun run build

build:
	make build-npm &&  go build cmd/main.go

run-prod:
	RUNENV=build ./main