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

## File Structure:

Generated with `tree`:

```

```

## Config

No Config yet.
