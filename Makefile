# go build command
build:
	@echo " >> building binaries"
	@go build -v -o autobase-twitter cmd/rest/*.go

# go run command
run: build
	./autobase-twitter

# run all go:generate commands (eg. Mock files generator)
generate:
	@go generate ./...