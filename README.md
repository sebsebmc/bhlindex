# Biodiversity Heritage Library Scientific Names Index

Creates an index of scientific names occurring in the collection of literature
in Biodiversity Heritage Library

## Usage

*NOTE*: This is a Dev release.

### Mac OSX

* Download [bhlindex release for mac][bhlindex-mac].
* Untar the file go to `script` directory and [read instructions][readme].
* Use [bhl testdata][bhl-test] for testing.

### Linux

* Download [bhlindex release for Linux][bhlindex-linux]
* Untar the file go to `script` directory and [read instructions][readme].
* Use [bhl testdata][bhl-test] for testing.

## Database Migrations

### Install migrate

```bash
go get -u -d github.com/golang-migrate/migrate/cli github.com/lib/pq
go build -tags 'postgres' -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/cli
```

### Create migration

```bash
migrate -ext sql -D db NAME
```

### Run commands

```bash
migrate -database postgres://localhost:5432/database up 2
```

### Commands

create [-ext E] [-dir D] NAME
: Create a set of timestamped up/down migrations titled NAME, in
  directory D with extension E

version
: current migration version

up [N]
: up N migrations

down [N]
: down N migrations

drop
: nuke database

### Testing

```bash
docker-compose build
docker-compose up
```

To update all dependencies change LAST_FULL_REBUILD line in Docker file and
return `docker-compose build`

## Importer
This fork at /sebsebmc/bhlindex adds a basic grpc import server in addition to
existing bhlindex grpc server. The import server allows json data to be sent
to the database and tied to pages or titles. Currently, there is a working 
example at /importer/example/ that just iterates over the pages and inserts
a counter as json data for that page.

### Running the example
I used a setup where I ran the docker container and used `bhlindex find` to
populate the database with pages. In another terminal I ran `bhlindex server`
to run the grpc server that streams pages. In the next terminal I ran 
`bhlindex importer` to run the import server. Finally, I used 
`go run test_client.go` in the /importer/example directory to run the test 
client.

The end result is a column in `page_imports` where the data is in the form of
`{"test": %d}" where %d is a running counter.

I made use of the `env` command in order to pass the environment values to 
bhlindex so that it would connect to the postgres server that was running
inside of docker.

[bhlindex-mac]: https://github.com/gnames/bhlindex/releases/download/v0.1.0/bhlindex-0.1.0-mac.tar.gz
[bhlindex-linux]: https://github.com/gnames/bhlindex/releases/download/v0.1.0/bhlindex-0.1.0-linux.tar.gz
[bhl-test]: https://github.com/gnames/bhlindex/releases/download/v0.1.0/bhl-testdata.tar.gz
[readme]: https://github.com/gnames/bhlindex/tree/master/bhlindex
