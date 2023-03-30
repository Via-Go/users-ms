build:
	@go build -o bin/ ./...

run: build
	./bin/src

test:
	@go test -v ./...

cert:
	cd cert && ./create.sh && cd ..

.PHONY: gen clean server client test cert