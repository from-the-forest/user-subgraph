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