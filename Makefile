# run gql codegen
generate:
	go run github.com/99designs/gqlgen generate

# sets up project
setup: install generate

# starts server with hot reloading
dev:
	npx nodemon \
		--exec go run ./server.go \
		--signal SIGTERM

# install go and npm dependencies
install:
	go mod download
	npm install

# builds executable
build: 
	go build -o ./bin/user-subgraph

# lint the schema
# NOTE: gql-schema-lint package doesn't support federation directives :(
lint:
	echo "not implemented"; exit 0;

# run unit tests
test:
	# TODO: arg for watch mode
	go test ./...

# run unit tests and capture coverage
cover:
	echo "not implemented"; exit 1;

# run integration tests
integration:
	ENV=staging echo "not implemented"; exit 1;

# starts server without hot reloading
start:
	go run ./server.go



