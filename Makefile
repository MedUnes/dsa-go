debug:
	dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --log
test:
	go run gotest.tools/gotestsum@latest --format=testdox