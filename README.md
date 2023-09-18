# Practice 

What am I practicing? I am using this repostory to experiment with different technologies that I may be unfamiliar with. 

For instance, I have not worked with the Go programing language yet, and I have used a little bit of GraphQL, so I made a simple GraphQL api, that uses a Go backend.

### List of languages/technologies that I have practiced with:
- Go
- GraphQL
- Docker
- Google Cloud Run
- Google Artifact Respository
- GoogleCloud CLI

## What I did
I looked at a few YouTube videos about Go and GraphQL. Then I created my own repository. It is a really simple API that, at the moment, only has a few functions. You are able to see what is possible by looking at the [GrpahQL schema](https://github.com/Whelans90/practice/blob/main/graph/schema.graphqls).

I used the Go mod GqlGen which generates code based on the GraphQl Schema. Then I created resolver functions, that modify an in memory datastore. Once I was done I created a Docker Image that I pushed to Google Artifact Repository, and deployed with Google Cloudrun. 

## In the Future
I would like to add a DB connection, maybe a PostgreSQL DB. I also need to add more api endpoints. Ontop of that I would like to improve the error handling 

## Where can you see the results?

This first draft has been [published on CloudRun](https://go-practice-graphql-gqszpzvufq-ul.a.run.app/), I am not sure how much this will cost in the future, but for now I have a free trial. 

### What to do with it?

Create some GQL queries. I have a mutation for seeding team data in the in memory data store. Right now it can only be run once without any validations, but in the future I would like to correct that. You can run it with the following command:

```graphql 
mutation seedTeams() {
    seedTeams(input: true) {
        name
        league
        id
    }
}
```
**Upsert a player**
```graphql 
mutation upsertPlayer($player: PlayerInput!) {
    upsertPlayer(input: $player) {
        code
        success
        message
        player{
            name
            id
            position
            team{
                    ame
            }
        }
    }
}
```
Sample JSON variable:
```json
{
  "player": {
    "name": "Kevin De Bruyne",
    "position": "MIDFIELDER",
    "nationality": "Belgium",
    "teamName": "MANCHESTERCITY"
  }
}
```

**Fetch a player**
```graphql 
query player($id: ID!) {
    player(id: $id) {
        name
        position
        id
        team {
            name
        }
    }
}
```
**Fetch Teams**
```graphql 
query teams {
  teams{
    name
    id
    league
  }
}
```

