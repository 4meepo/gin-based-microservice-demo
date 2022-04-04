./cmd/wire_gen.go: ./cmd/wire.go
	cd cmd && wire
build:
	go build -o app ./cmd
test:
	go test ./...