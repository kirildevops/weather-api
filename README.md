# Multi methods

## Instructions

0. Check the `app.env` file. It contains the config of our app.
    This information is sensitive and commited to the repo just for the development purposes.  
    `Viper` will read this file and populate some configs.

Run the next set of commands to check the work
1. `make postgres` to create the postgres container in the docker network. [TBD use Docker Compose to spin up the DB]
2. `make migrateup` to run the migrations
3. `make sqlc` to generate sqlc idiomatic Go code to interact with Postgresql
4. `make test` to run DB tests [only weather-api/db/sqlc is covered at this time]
5. `make server` to run the server [only `subscribe` endpoint is functional as of 22:00 15May2025]

6. Open the new terminal window and run `curl -X POST -H "Content-Type: application/json" -d @req1.json http://localhost:8080/api/subscribe | jq`

7. Connect to the db via DBeaver or alike tool to check the result. There are some remnants of testing with garbage data. It could be cleaned up later.

8. Use `make proto` to build the `pb.go` files in the `pb` directory [UNDER CONSTRUCTION]
