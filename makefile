dev-go:
	go run cmd/main.go

dev-npm:
	cd App && npm run start

dev:
	make -j 2 dev-npm dev-go