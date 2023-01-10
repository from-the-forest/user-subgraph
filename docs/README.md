# User Subgraph Documentation

## General Principles

- Easy to develop, debug, test, and deploy quickly.  It's essential to be able to iterate quickly.
- Reliable deployments. Developers should never be afraid to break things - proper checks should be in place to prevent that from happening
- Compute agnostic.

## GraphQL Principles

- The schema is the source of truth.
- It's up to the developer / team to decide if logic should live in resolvers or in downstream services.  That is a decision best made based on scaling and compute requirements.


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
  - [ ] data loader(s) in context / database (dockerized and non dockerized)
  - [ ] Parse request context to get user
- [ ] Unit Testing
  - [x] 1st test written
  - [ ] Coverage Report
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
  - [ ] Primary Workflow
    - [x] Install
    - [x] Secrets
    - [x] Build
    - [x] Schema Linting
    - [x] Unit Testing
    - [x] Infra
    - [x] Deployment
    - [ ] Publish schema to registry
    - [ ] Integration Testing
  - [ ] Nightly Workflow
    - [x] tear down stack (to avoid extra costs)
    - [ ] check for packages that are out of date
- [ ] Deployment
  - [x] Dockerize
  - [ ] deploy in a compute agnostic way
    - [ ] kubernetes
    - [ ] serverless
    - [ ] standalone

