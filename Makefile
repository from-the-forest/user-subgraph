# builds executable
build: 
	go build

# lint the schema / go files
lint:
	echo "not implemented"; exit 1;

# run unit tests
test:
	echo "not implemented"; exit 1;

# run integration tests
integration:
	echo "not implemented"; exit 1;

# starts server without hot reloading
start:
	go run ./server.go

# starts server with hot reloading
dev:
	npx nodemon \
		--exec go run ./server.go \
		--signal SIGTERM

# run codegen
generate:
	go run github.com/99designs/gqlgen generate