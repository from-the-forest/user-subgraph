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
  users(input: { first: 5 }) {
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
  node(id: "VXNlcjpjMGQzODQxNC1hYjE5LTQwODAtYTZiMi1jNWFkNjEyMTYwZjQ=") {
    id
    ... on User {
      ...UserFields
    }
  }
}

mutation CreateUser {
  createUser(input: { firstName:"Willy", lastName: "Cuffney", email: "littlewilly@dog.com" }) {
    ...UserFields
  }
}

mutation UpdateUser {
  updateUser(input: { firstName:"Nox", lastName: "Cuffney", email: "noxman@dog.com" }) {
    ...UserFields
  }
}

mutation DeleteUser {
  deleteUser(id: "VXNlcjpkODBhOTNiZS00MGEwLTRhNTctODQ2YS1lZTU5MDY1ZmY1Mzc=") {
    ... UserFields
  }
}
```

NOTE: you can only use @defer with the supergraph - not the subgrpah