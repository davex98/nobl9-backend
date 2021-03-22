![example workflow](https://github.com/davex98/nobl9-backend/actions/workflows/tests.yml/badge.svg)

# Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [How to](#how-to)
    * [Run the app without building it](#run-app-without-building-it)
    * [Run the development environment](#run-the-development-environment)
    * [Run the unit tests](#run-the-unit-tests)
    * [Reload the api specification](#reload-the-api-specification)
    * [Run the end to end tests](#run-the-end-to-end-tests)
- [API specification](#api-specification)


# Overview

This repository is a solution for a Golang Developer role in [Nobl9](https://nobl9.com/).

It's an implementation of a service that gets random numbers from [random.org](https://random.org/) and calculates standard deviation of the drawn integers and additionally standard deviation of sum of all sets.
# Project Structure

- [api](api/) - OpenAPI definitions
- [docker](docker/) - Dockerfiles, one for production (optimized one **~8Mb**, running the application as nonroot) and one for development with live-reload 
- [cypress/integration](cypress/integration/) - e2e tests specification
- [random-generator](random-generator/) - application code, it is written using clean architecture approach

# Prerequisites
Make sure you have installed all of the following prerequisites on your development machine:
* Git - [Download & Install Git](https://git-scm.com/downloads). OSX and Linux machines typically have this already installed.
* Node.js - [Download & Install Node.js](https://nodejs.org/en/download/) and the npm package manager. If you encounter any problems, you can also use this [GitHub Gist](https://gist.github.com/isaacs/579814) to install Node.js.
* Docker - [Download & Install Docker](https://docs.docker.com/engine/install/ubuntu/). Docker is used for building images and running the end-to-end tests.
* Golang - [Download & Install Golang](https://golang.org/doc/install).

The applications requires setting up 2 environment variables (when running from [docker image](https://hub.docker.com/repository/docker/jakuburghardt/nobl9-backend), they are set to default values):
- **PORT=8080**
- **CONCURRENT_REQUESTS=5** - Limit of concurrent requests
  
*As the provided date range in a single request might be broad, the random API should be queried
  concurrently. However, in order not to be recognized as a malicious user, a limit of concurrent
  requests to this external API must exist.*


# How to

## Run app without building it
In order to run the service you can simply run:

```bash
make run_app_from_repository
```
It will automatically get the image from the repository and expose the port on **8080**.

## Run the development environment
In order to run the development environment, that enables the auto-reload so that we can see live changes, run:
```bash
make dev_env
```

## Run the unit tests
In order to run the unit tests, run:

```bash
make unit_test
```


## Reload the api specification
Once the api specification got changed in [api/swagger.yml](api/swagger.yml) we need to update the code generated from that template. In order to do that, run:

```bash
make openapi
```

## Run the end to end tests
In order to run the end-to-end tests using cypress, run:

```bash
make e2e_test
```


# API specification
[SwaggerHub](https://app.swaggerhub.com/apis/burghardtjakub/Nobl9-backend/1.0.0)
