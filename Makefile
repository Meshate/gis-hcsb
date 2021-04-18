all: gotool
clean:
	rm -f hcsb
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	gofmt -w .
	go vet -v .
	CGO_ENABLED=0 go build -o hcsb -ldflags "-w -s" -v ./cmd/main.go
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"
.PHONY: clean gotool help