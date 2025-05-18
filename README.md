# Overview

## Solution Stack

0. Project structure have been influenced by following along with the course https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes Some parts were replicated (NOT copied) from the course video.

1. SQL database. Postgres 12 in this case. Potentially can be replaced with more recent version. Runs inside Docker
2. `sqlc` to interact with Postgres. And the `lib/pq` as the driver. Can be replaced with `pgx` later. (SQLC docs)[https://docs.sqlc.dev/en/stable/guides/using-go-and-pgx.html]
3. `golang-migrate` to facilitate DB migrations. Currently as a standalone CLI. Can be moved to the container version.
4. `viper` to load configs.
5. `dotenv` to load the API_KEY for the 3rd party service.
6. `testify` to run asserts, mocks etc.
7. `uuid` from Google to generate random unique string that we use as a `token`
8. `gin-gonic` as a web-framework to serve our API
9. `net/http` standard package to query external API
10. `protoc` compiler to work with protobufs [not finished yet]
11. `Makefile` to make some aliasing simpler.

## Workflow Instructions

0. Check the `app.env` file. It contains the config of our app.
    This information is sensitive and commited to the repo just for the development purposes.  
    `Viper` will read this file and populate some configs.  
    Put your Weather API key into `.env` file in the root of the repository. Precede it with `WEATHER_API_KEY=`

Run the next set of commands to check the work
1. `make postgres` to create the postgres container in the docker network. [TBD use Docker Compose to spin up the DB]
2. `make migrateup` to run the migrations.
3. `make sqlc` to generate sqlc idiomatic Go code to interact with Postgresql
4. `make test` to run DB tests [only weather-api/db/sqlc is covered at this time]
5. `make server` to run the server

6. Test **Subscribe** Endpoint: Open the new terminal window and run `curl -X POST -H "Content-Type: application/json" -d @req1.json http://localhost:8080/api/subscribe | jq`

7. Connect to the db via DBeaver or alike tool to check the result. There are some remnants of testing with garbage data. It could be cleaned up later.

8. Test **Confirm** Endpoint: Open the new terminal window and run `curl -X GET http://localhost:8080/api/confirm/2d77cd87-161f-4efb-a9f0-8831aa52cd44` 

9. Test **Weather** Endpoint: Open the new terminal window and run `curl -X GET http://localhost:8080/api/weather?city=Kyiv`

10. Test **UnSubscribe** Endpoint: Open the new terminal window and run `curl -X GET http://localhost:8080/api/unsubscribe/3d64fede-e088-4b15-9b4c-94cf1` where the last part is the token which can be copied from the DB column.

100. Use `make proto` to build the `pb.go` files from `proto` into the `pb` directory [UNDER CONSTRUCTION]

