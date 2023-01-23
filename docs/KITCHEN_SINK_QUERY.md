```gql
fragment UserFields on User {
  id
  firstName
  lastName
  email
}

query KitchenSink {
  whoami {
    ...UserFields
  }
  _service {
    sdl
  }
}
```

NOTE: you can only use @defer with the supergraph - not the subgrpah