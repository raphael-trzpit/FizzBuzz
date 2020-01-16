# FizzBuzz
FizzBuzz Go server

## Install

`cd cmd/api && PORT=8080 go run .`

config from environment variables:
 - PORT (http port)

## Limitations

The repository for the statistics is in memory, meaning we can't scale horizontally. We just need to implement a new repo with a database to store a statistics if we want to make it scallable.

Some errors / logs might not be "clean", but i tried not to use any library

HTTP Method are not checked or used in the routing (using standard router instead of mux or httpdrouter)

