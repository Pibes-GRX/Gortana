BINARY_NAME=prueba

test:
	go test tests/aserciones_test.go

build:
	echo "Compilo $(BINARY_NAME)"
	go build -o bin/$(BINARY_NAME) src/$(BINARY_NAME).go

run: build
	./bin/$(BINARY_NAME)

go:
	go run src/$(BINARY_NAME).go

clean:
	rm -rf bin/*
	rm logPibes.txt