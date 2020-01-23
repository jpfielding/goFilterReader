all: restore-deps test vet

test:
	go test -v ./filter
vet: 
	go vet ./filter
clean:
	rm *.test
restore-deps:
	go mod tidy