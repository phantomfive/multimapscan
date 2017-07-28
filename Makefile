scanner: main.go
	go build

run: scanner
	time ./scanner
