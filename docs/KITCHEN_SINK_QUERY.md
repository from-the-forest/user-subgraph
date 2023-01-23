```gql
fragment UserFields on User {
      id
    firstName
    lastName
    fullName
    email
}

query KitchenSink {
  whoami {
		...UserFields
  }
  users {
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    edges {
      node {
        ...UserFields
      }
    }
  }
  node(id: "VXNlcjox") {
    id
    ... on User {
      ...UserFields
    }
  }
  host
}

# mutation DeleteUser {
#   deleteUser(id: "VXNlcjox")
# }
```

NOTE: you can only use @defer with the supergraph - not the subgrpah