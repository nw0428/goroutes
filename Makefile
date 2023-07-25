


build:
	go get . && go build example_server.go rps.go

run: build
	./example_server