# run gql codegen
generate:
	go run github.com/99designs/gqlgen generate

# sets up project
setup: env install generate

# starts server with hot reloading
dev:
	npx nodemon \
		--exec go run ./server.go \
		--signal SIGTERM

# create template .env file
env:
	cp .env.sample .env

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
	go test ./graph/...

# run unit tests and capture coverage
cover:
	echo "not implemented"; exit 1;

# run tests tests
integration:
	GRAPHQL_ENDPOINT="https://main--cuffney-supergraph-vokiem.apollographos.net/graphql" go test ./tests/integration/...

# run tests tests
e2e:
	go test ./tests/e2e/...

# starts server without hot reloading
start:
	go run ./server.go

docker-build:
	docker build -t ftf/user-subgraph:latest --build-arg MONGO_CONNECTION_STRING="mongodb+srv://doadmin:fl79j0845VD6Sv1Q@ftf-main-cluster-20fa74e3.mongo.ondigitalocean.com/admin?authSource=admin&replicaSet=ftf-main-cluster&tls=true" .

docker-run:
	docker run -p 4000:4000 ftf/user-subgraph:latest

