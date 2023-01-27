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
  searchUsers(first: 2, after: "VXNlcjo3MDllOGU1Yy02ODVjLTRiMDUtYjNhNy1hNTJkYmY3ZGMxZDM=") {
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
  node(id: "VXNlcjowODEyMzE3ZC1hYzE0LTRkODktOTMwZi03MDgyMmZjNzdjMGI=") {
    id
    ... on User {
      ...UserFields
    }
  }
  nodes(ids: ["VXNlcjowODEyMzE3ZC1hYzE0LTRkODktOTMwZi03MDgyMmZjNzdjMGI="]) {
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
  updateUser(input: { id: "VXNlcjo5YTIyNDZkNy04NGQ4LTQ5YzItOWQ3OS05MDU3ZjYxN2IyM2Y=", firstName:"Willow", lastName: "Cuffney", email: "noxman@dog.com" }) {
    ...UserFields
  }
}

mutation DeleteUser {
  deleteUser(id: "VXNlcjpkOWM0ODY1ZC00Yjg5LTRhMjEtYWVjNi1jZjAzM2I3M2E3N2M=") {
    ... UserFields
  }
}
```

NOTE: you can only use @defer with the supergraph - not the subgrpah