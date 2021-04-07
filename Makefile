.PHONY: all 
all:  
# go test -v ./...
	go test -v -args -n 22

bench:
	go test -benchmem -bench . -args -n 32
	go test -benchmem -bench . -args -n 1000
	go test -benchmem -bench . -args -n 1000000