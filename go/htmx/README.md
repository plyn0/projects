# install Go
https://go.dev/doc/install

# setup VSCode
To get type linting, auto-format, auto-imports, etc:\
https://marketplace.visualstudio.com/items?itemName=golang.Go
Add the templ extension to lint template files.

# how the project was started
- `go mod init poc-go` creates the `go.mod` file (necessary to have dependencies).
- `go get -u github.com/gin-gonic/gin` to add the GIN framework (creates the `go.sum` file).
- `github.com/stretchr/testify` for the tests (https://gin-gonic.com/docs/testing/).

# install dependencies after cloning project
`go mod tidy` to install the dependencies specified in the `go.sum` file (but `go run main.go` can also take care of it before running the code).

# generate go code from templ files
`templ generate`

# running the project
- `go run main.go` to run the project.
<!-- - `go test` to run the tests. -->

## live reload in development
https://github.com/air-verse/air install air utility.
`air` in the root folder.
`templ generate -watch` in another terminal so that templating changes are also part of the live reload. 

<!-- # running tests
`go test ./...` to run the tests.
`go test ./... -count=1` to run the tests without caching previous results.

# run with Docker
`docker build . -t poc-go` to build the image.\
`docker container run -p 8080:8080 poc-go` to run the project. -->

# project structure
`public/index.html` contains the root of the website (static content). When user actions occur (button, form), requests are issued to the GIN routes (main.go). Those return HTML (templates/ ?) by using `templ` to inject parameters (index.templ ?).