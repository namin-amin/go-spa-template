dev-go:
	go run cmd/main.go

dev-npm:
	cd ui && npm install && npm run start

dev:
	make -j 2 dev-go dev-npm

build-npm:
	cd ui && npm install && npm run build

build:
	make build-npm &&  go build cmd/main.go

run-prod:
	./main