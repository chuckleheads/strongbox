# Strongbox

Yet another attempt at a simple Habitat depot. Right now, this only supports
origin creation and retrieval.

## Prerequisites
Postgres must be running and have the `strongbox` db created, with an `origins`
table. There are no migrations here. Here's some sample SQL for the lazy:

```sql
CREATE TABLE origins (id bigserial PRIMARY KEY, name text NOT NULL UNIQUE);
```

Install the appropriate deps:
```shell
go get -u github.com/go-pg/pg
go get -u github.com/labstack/echo
```

## Setup
`go run main.go`

## Example requests

Creating an origin:
`curl -v -H "Accept: application/json" -H "Content-Type: application/json" -d '{"name":"core"}' http://localhost:1323/origins`

Retrieving an origin:
`curl -v http://localhost:1323/origins/core`
