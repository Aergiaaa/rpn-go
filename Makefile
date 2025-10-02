build:
	@go build . -o

build-win:
	@GOOS=windows GOARCH=amd64 go build .

run: build
	@./rpn-go