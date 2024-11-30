# install Go
https://go.dev/doc/install

# setup VSCode
To get type linting, auto-format, auto-imports, etc:\
https://marketplace.visualstudio.com/items?itemName=golang.Go

# how the project was started
- `go mod init poc-go` creates the `go.mod` file (necessary to have dependencies).
- `go get -u github.com/gin-gonic/gin` to add the GIN framework (creates the `go.sum` file).
- `github.com/stretchr/testify` for the tests (https://gin-gonic.com/docs/testing/).

# install dependencies after cloning project
`go mod tidy` to install the dependencies specified in the `go.sum` file (but `go run main.go` can also take care of it before running the code).

# running the project
- `go run main.go` to run the project.
- `go test` to run the tests.

## live reload in development
https://github.com/air-verse/air install air utility.
`air` in the root folder.

# running tests
`go test ./...` to run the tests.
`go test ./... -count=1` to run the tests without caching previous results.

# run with Docker
`docker build . -t poc-go` to build the image.\
`docker container run -p 8080:8080 poc-go` to run the project.

# calls scenario
## get products
- `curl 'http://localhost:8080/api/v1/products'` returns all 20 products.
- `curl 'http://localhost:8080/api/v1/products?category=fruits'` returns 10 products.
- `curl 'http://localhost:8080/api/v1/products?category=fruits&sort=price&offset=2&limit=4'` returns 4 products, with prices from `10.5` to `19.9`.

## put cart
- `curl -X PUT 'http://localhost:8080/api/v1/carts/1' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"productId": 2, "quantity": 3}'` should return `{"cartId":1,"products":[{"productId":1,"quantity":3},{"productId":2,"quantity":5},{"productId":11,"quantity":1}]}` on the first request.
- `curl -X PUT 'http://localhost:8080/api/v1/carts/2' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"productId": 3, "quantity": 1}'` to add a product that was not in the cart. Should return `{"cartId":2,"products":[{"productId":7,"quantity":2},{"productId":13,"quantity":3},{"productId":5,"quantity":4},{"productId":3,"quantity":1}]}` on the first request.

## get cart
- `curl 'http://localhost:8080/api/v1/carts/1'` to get a cart.
- `curl 'http://localhost:8080/api/v1/carts/wrong'` checks that the id is not an int.
- `curl 'http://localhost:8080/api/v1/carts/999'` checks that the requested cart does not exist.

## concurrent call
- `curl 'http://localhost:8080/api/v1/concurrent'` to request the concurrent endpoint
