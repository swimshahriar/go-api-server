### Getting started

You need to have `go` installed in your machine to run this app. Execute `go run cmd/main.go` to start the server on port `3000`.

### Available routes

- `/healthcheck` - GET - Returns if server is working or not
- `/api/products` - GET - Returns all the products
- `/api/products` - POST - Add products (pass `title` in the req body)
