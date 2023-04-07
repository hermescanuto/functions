

build:
	go get ./...
	go build cmd/api.go

run-verbose:
	func start --verbose

run:
	func start 