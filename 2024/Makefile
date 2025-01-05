PACKAGE=./...

build: lint
	go build

run: build
	./2024

test:
	go test ${PACKAGE}

benchmark:
	go test ${PACKAGE} -bench=. -benchtime=5s

clean:
	go clean

lint: format
	golangci-lint run ./... --fast --config=./.golangci.yaml

format:
	gofumpt -l -w .
