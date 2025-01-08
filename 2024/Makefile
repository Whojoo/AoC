PACKAGE=./...

build: lint
	go build

run: build
	./2024

test:
	go test ${PACKAGE}

benchmark:
	go test ${PACKAGE} -bench=.

benchmark-stats:
	go test ${PACKAGE} -bench=. -count=20 -benchmem | tee stats.txt
	benchstat stats.txt

clean:
	go clean
	go mod tidy
	rm stats.txt

lint: format
	golangci-lint run ./... --fast --config=./.golangci.yaml

format:
	gofumpt -l -w .
