# go build command
build-worker:
	@echo " >> building worker binaries"
	@go build -v -o autobase-twitter-worker cmd/worker/*.go

# go run command
run-worker: build-worker
			./autobase-twitter-worker

# go build command
build-api:
	@echo " >> building api binaries"
	@go build -v -o autobase-twitter-api cmd/rest/*.go

# go run command
run-api: build-api
		./autobase-twitter-api

# run all go:generate commands (eg. Mock files generator)
generate:
	@go generate ./...