# Clean Architecture Simple App

[![Build](https://github.com/cyruzin/clean_architecture/workflows/Build/badge.svg)](https://github.com/cyruzin/clean_architecture/actions?query=workflow%3ABuild+branch%3Amaster) [![Go Report Card](https://goreportcard.com/badge/github.com/cyruzin/clean_architecture)](https://goreportcard.com/report/github.com/cyruzin/clean_architecture) [![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)

This repo contains a simple command line tool and a rest server that checks the best price between two destinations.

## Architecture

For the system architecture I have decided to use the Clean Architecture proposed by Robert C. Martin (Uncle Bob).

Pull requests are welcome.

<p align="center"><img src="./assets/clean_architecture_diagram.png"></p>

## Install

Make sure you have Git and Go installed in you machine.

Clone the repo:

```sh
 git clone git@github.com:cyruzin/clean_architecture.git
```

Install Go dependencies:

```go
 go mod download
```

If you want to modify the default values, rename the .env.example file to .env and change the values if you want.

```sh
 mv .env.example .env
```

## Running

### Cli

Go to the cli folder:

```sh
 cd cmd/cli
```

Then, run the command below:

```go
  go run main.go BBB AAA
```

### Rest Server

Go to the routes folder:

```sh
 cd cmd/routes
```

Then, run the command below:

```go
  go run main.go
```

Default Base URL: http://localhost:8000

End-points:

GET - /route

Params required: departure, destination

Eg: http://localhost:8000/route?departure=BBB&destination=AAA

POST - /route

Params required: departure, destination, price

Eg: http://localhost:8000/route

JSON Payload:

```json
{
  "departure": "BBB",
  "destination": "AAA",
  "price": 76
}
```

## Libs

- go-chi
- json-iterator
- zerolog
- envconfig

## License

MIT
