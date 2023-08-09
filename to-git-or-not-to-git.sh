curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"jesaed\"}}){id}}"}' | jq ' .data.user[0].id'

