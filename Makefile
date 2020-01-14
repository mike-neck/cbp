.PHONY: build
build:
	go build -o build/cbp -v

.PHONY: clean
clean:
	go clean
	rm -rf build/
