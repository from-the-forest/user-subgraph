# User Subgraph

## Requirements

- `go`
- `npm` & `node` (only needed for (hot reloading) during local development)

## Getting Started

1. `go mod install`
2. `make dev`
3. `open http://localhost:8000`

## Tasks

- [x] use `gin` as a server rather than `net/http` to allow for middleware (custom context)
- [ ] DevEx
  - [x] Hot Reloading
  - [ ] Debugging
- [x] Secrets
  - [x] `.env` support
- [x] Context
  - [x] env in context
  - [x] user comes from context
  - [ ] data loader(s) in context
  - [ ] Parse request context to get user
- [ ] Unit Testing
- [ ] Schema
  - [ ] Scalars
- [ ] Relay
  - [ ] Pagination Utils
  - [ ] Node / Nodes query
- [ ] Federation
  - [ ] Entities
  - [ ] Field Resolvers
- [ ] Directives
  - [ ] @depricated
- [x] CI
  - [x] Install
  - [x] Secrets
  - [x] Build
  - [x] Schema Linting
  - [x] Unit Testing
  - [ ] Integration Testing
- [ ] Deployment
  - [ ] Dockerize
  - [ ] deploy as
    - [ ] serverless
    - [ ] standalone
    - [ ] kubernetes