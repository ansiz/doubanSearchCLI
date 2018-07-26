build:
	go build -o dsearch cmd/cli.go
install:
	make build
	mv dsearch /usr/local/bin/
.PHONY: clean
clean:
	-rm dsearch
