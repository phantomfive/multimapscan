multimapscan: main.go
	go build

run: multimapscan
	time ./multimapscan
