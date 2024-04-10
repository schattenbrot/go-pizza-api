# Go Pizza API

This is a pizza API written in Go.

## Create Swagger

To create swagger documentation run:

> docker run --rm -v $(pwd):/code ghcr.io/swaggo/swag:latest init -g ./cmd/api/main.go

The documentation is available on url:

> http://localhost:8080/docs

> http://localhost:8080/docs/doc.json

## API Functions

Base URL:

> {domain:port}/api/{version}/{endpoint}

Currently there only exists `version` `v1`.

| Endpoint        | Method | Description                                |
| --------------- | ------ | ------------------------------------------ |
| `/` & `/status` | `GET`  | Both show the api status                   |
| `/ping`         | `GET`  | Returns a success response with no content |

# Status Codes

For simplicity this api only uses a couple status codes:

-   200 OK
-   201 Created
-   204 No Content
-   400 Bad Request
    -   conversion from string to ObjectID will also fall here even thought it might actually be a 422 Bad Data Error instead
-   404 Not Found
-   500 Internal Server Error

## File Structure:

Generated with `tree`:

```
.
├── cmd
│   └── api
│       ├── main.go
│       └── models.go
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── auth
│   ├── config
│   │   └── config.go
│   └── services
│       ├── app
│       │   ├── controllers.go
│       │   └── routes.go
│       ├── orders
│       │   ├── controllers.go
│       │   ├── db.go
│       │   ├── middlewares.go
│       │   ├── models.go
│       │   └── routes.go
│       └── pizzas
│           ├── controllers.go
│           ├── db.go
│           ├── middlewares.go
│           ├── models.go
│           └── routes.go
├── pkgs
│   └── responder
│       └── responder.go
└── README.md
```

## Config

No Config yet.

## TODO

-   a log ...
-   don't forget auth D:
