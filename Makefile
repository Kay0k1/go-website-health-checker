run:
	go run main.go

build:
	go build -o bin/checker main.go

test:
	go test -v ./...

clean:
	rm -rf bin/