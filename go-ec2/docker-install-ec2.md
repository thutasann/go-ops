# Install Docker on an Ubuntu EC2 Instance

This guide installs Docker on the Ubuntu EC2 instance used for the Go app. Run these commands in your EC2 terminal while logged in as `ubuntu`.

The instance shown in the screenshot runs Ubuntu 26.04 LTS. These instructions use Ubuntu's `docker.io` package. [Ubuntu package details](https://packages.ubuntu.com/en/resolute-updates/docker.io)

## Install and start Docker

Run each command in order:

```bash
sudo apt-get update
sudo apt-get install -y docker.io
sudo systemctl enable --now docker
sudo usermod -aG docker ubuntu
```

| Command                         | Purpose                                                                                 |
| ------------------------------- | --------------------------------------------------------------------------------------- |
| `apt-get update`                | Refresh the available package list.                                                     |
| `apt-get install -y docker.io`  | Install Docker using Ubuntu's package repository.                                       |
| `systemctl enable --now docker` | Start Docker and enable it to start at boot.                                            |
| `usermod -aG docker ubuntu`     | Add the `ubuntu` user to the Docker group so it can run Docker commands without `sudo`. |

The Docker group grants privileges equivalent to root. Avoid the screenshot's `chmod 666 /var/run/docker.sock` command, which gives every local user access to the Docker socket. [Docker post-installation instructions](https://docs.docker.com/engine/install/linux-postinstall/)

## Apply your group membership

Disconnect from the EC2 terminal and reconnect. Your new session will pick up the Docker group membership.

Alternatively, start a shell with the updated group in your current session:

```bash
newgrp docker
```

## Verify the installation

```bash
docker --version
docker run --rm hello-world
docker ps
```

Expected results:

- `docker --version` prints the installed version.
- `docker run --rm hello-world` downloads and runs a test container, prints **Hello from Docker!**, and removes the container after it exits. This requires outbound internet access.
- `docker ps` lists running containers. An empty list is normal after the test container exits.

For daemon status, run:

```bash
sudo systemctl status docker --no-pager
```

Look for `active (running)`. These verification commands confirm Docker is ready; the Go application still needs to be built and run separately.

If a Docker command reports permission denied, reconnect after adding the user to the group, then check membership with `id -nG`. The output should include `docker`.
