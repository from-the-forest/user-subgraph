# User Subgraph

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/from-the-forest/user-subgraph/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/from-the-forest/user-subgraph/tree/main)

## Requirements

- `go`
- `npm` & `node` (only needed for (hot reloading) during local development)

## Getting Started

1. `make setup`
2. ensure your `.env` file is completely filled out
3. `make dev`
4. `open http://localhost:4000`

## Commands

`$ npm run <command>`

| command                     | description   |
|-----------------------------|---------------|
| `make install`              | install go and npm dependencies |
| `make build`                | build a binary of the application |
| `make lint`                 | lints the schema |
| `make test`                 | runs unit tests once |
| `make integration`          | runs integration tests against specified env |
| `make cover`                | gather test coverage |
| `make start`                | starts the subgraph's production server |
| `make dev`                  | starts the subgraph's dev server |
| `make generate`             | runs gqlgen code generation |

# Deployments

CI is setup to do tag based deployments.

#### Create a tag

`git tag -a v0.1.0 -m "initial deployment"`

#### View tags

`git tag`

#### Push tag

`git push origin v0.1.0`

