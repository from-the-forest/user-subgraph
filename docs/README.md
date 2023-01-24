# User Subgraph Documentation

## General Principles

- Easy to develop, debug, test, and deploy quickly.  It's essential to be able to iterate quickly.
- Reliable deployments. Developers should never be afraid to break things - proper checks should be in place to prevent that from happening
- Compute agnostic.

## GraphQL Principles

- The schema is the source of truth.
- It's up to the developer / team to decide if logic should live in resolvers or in downstream services.  That is a decision best made based on scaling and compute requirements.

## K8s help

https://www.youtube.com/watch?v=cJKdo-glRD0
https://www.youtube.com/watch?v=MpovOI5eK58

## MongoDB

- best practice to have 1 "cluster" with multiple "databases" and "collections" per subgraph. however that would be super expensive - so maybe 1 cluster for now.