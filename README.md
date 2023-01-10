# User Subgraph

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/jcuffney/user-subgraph/tree/main.svg?style=svg&circle-token=bbbd0224a449733c353d1454e72ee1982c01d7a2)](https://dl.circleci.com/status-badge/redirect/gh/jcuffney/user-subgraph/tree/main)
[![Asana](https://img.shields.io/badge/asana-%23273347.svg?&style=for-the-badge&logo=asana&logoColor=white)](https://app.asana.com/0/1203689378138415)

## Requirements

- `go`
- `npm` & `node` (only needed for (hot reloading) during local development)

## Getting Started

1. `make setup`
2. `make dev`
3. `open http://localhost:8000`

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
