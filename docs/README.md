# User Subgraph Documentation

## General Principles

- Easy to develop, debug, test, and deploy quickly.  It's essential to be able to iterate quickly.
- Reliable deployments. Developers should never be afraid to break things - proper checks should be in place to prevent that from happening
- Compute agnostic.

## GraphQL Principles

- The schema is the source of truth.
- It's up to the developer / team to decide if logic should live in resolvers or in downstream services.  That is a decision best made based on scaling and compute requirements.
- The files on a given type do NOT always have to map 1:1 to a database's table.  In fact often it's better that they don't!

## GQLGen/GoLang Notes
- you need to infer types from context
- even though multiple graphql files are supported - if you need to split up your schema - it might be too complicated.  Consider a new subgraph?

## K8s help

https://www.youtube.com/watch?v=cJKdo-glRD0
https://www.youtube.com/watch?v=MpovOI5eK58

## MongoDB

- best practice to have 1 "cluster" with multiple "databases" and "collections" per subgraph. however that would be super expensive - so maybe 1 cluster for now.