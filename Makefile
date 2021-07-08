test:
	go test -count=1 -failfast ./...

cov:
	go test -race -coverprofile=coverage.out -covermode=atomic

build:
	go build -v .

dep:
	go get -v -t -d ./...