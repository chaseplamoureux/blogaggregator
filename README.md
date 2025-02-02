# blogaggregator

blogaggregator is a guided project from [boot.dev](https://www.boot.dev). Written in Go and retrieves. RSS feeds and stores them in local PostgreSQL database.

## Prerequisites

goose package for go

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

PostgreSQL : [Docs](https://www.postgresql.org/)

gatorconfig file

```bash
touch ~./gatorconfig.json
```
config file content:
```json
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```



## Usage

```bash

# register user
./blogaggregator register chase

# login
./blogaggregator login chase

# get list of commands
./blogaggregator help
```
