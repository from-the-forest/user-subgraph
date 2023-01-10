# install go and npm dependencies
install:
	go mod download
	npm install

# builds executable
build: 
	go build -o ./bin/user-subgraph

# lint the schema / go files
lint:
	npx graphql-schema-linter --except relay-page-info-spec

# run unit tests
test:
	go test ./...

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

docker-build:
	docker build -t ftf/user-subgraph:latest .

docker-run:
	docker run -it --rm --name user-subgraph ftf/user-subgraph:latest