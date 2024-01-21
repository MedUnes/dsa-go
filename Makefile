debug:
	dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --log
update:
	$(MAKE) build
	./bin/dsa
build:
	go build -ldflags="-s -w" -o ./bin/dsa
run:
	go run main.go
test:
	go run gotest.tools/gotestsum@latest --format=testdox