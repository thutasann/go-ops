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

```bash
docker ps

docker exec -it <container_id> /bin/sh

ls

ls -ahl

ldd server
```

**Scratch**

```bash
docker build -t test-server-scratch -f Dockerfile.scratch .
```
