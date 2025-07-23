# Related Docs

## Cross Compiling

```bash
docker run --rm -it alpine
```

```bash
apk add go

mkdir app

cd app

vi main.go

go build -o main main.go

./main
```

## Check

```bash
# ls -ahl main is a shell command that lists the contents of the directory named main with specific options:
ls -ahl main

# ldd main is a command used to show the shared libraries that the executable file named main depends on.
ldd main
```

## Compile

```bash
CGO_ENABLED=0 go build -o main-nocgo main.go

GOOS=darwin ARCH=amd64 go build -o main-darwin64 main.go
```

## Tools Check

```bash
go tool dist list
```
