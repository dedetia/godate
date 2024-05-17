# Godate Service

This service is a simple dating app project created with a simple hexagonal architecture and written using Golang.

## Contents

* [Structure Explanation](#structure-explanation)
* [Usage](#usage)
    * [Running App](#running-app)
    * [Testing](#testing)

## Structure Explanation

- `assets`: This directory contains static files such as images, css, etc.
- `build`: This directory is used to store build of the project.
- `config`: This directory contains application configuration such as configuration files, environment variables, etc.
- `internal`: This directory contains internal code that should not be accessed by other projects. It includes adapters
  for external services and the core of the application, which is the main business logic.
- `pkg`: This directory contains external code that can be used by other projects.
- `shared`: This directory contains code that is shared across multiple parts of the application.

## Used technologies

* [Go](https://go.dev/)
* [MongoDB](https://www.mongodb.com/)
* [Docker](https://www.docker.com/)

## Usage

### Running App

Here is the steps to run app with `docker compose`

```bash
# move to directory
$ cd workspace

# clone into YOUR $GOPATH/src
$ git clone https://github.com/dedetia/godate.git

# move to project
$ cd godate

# build & run the application
$ make run

# check if the containers are running
$ docker ps

# stop
$ make stop
```

### Testing

Use the following script to run unit tests

```bash
make test
```