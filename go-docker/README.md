# Simple Demo of Golang + Docker

## Commands

```bash
docker build -t test-server .
```

```bash
docker images |grep test-server
```

```bash
docker run --rm -it -p 8080:8080 test-server

curl localhost:8080
```
