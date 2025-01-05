PACKAGE=./...

build:
	go build

run: build
	./2024

test:
	go test ${PACKAGE}

benchmark:
	go test ${PACKAGE} -bench=. -benchtime=5s

clean:
	rm 2024
