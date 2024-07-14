# Hello-world
## Initialize the go module
```bash
go mod init learn-go-with-tests
```

## Install go packages
```bash
go install golang.org/x/tools/cmd/godoc@latest
go install github.com/rakyll/gotest@latest
```

## Run go test files
```bash
gotest ./...
go test ./...
```

## Enable git hooks scripts.
```bash
git config --local core.hooksPath .githooks
```

## Check workflows status
```
$ gh workflow view Lint
$ gh workflow view Test
```
