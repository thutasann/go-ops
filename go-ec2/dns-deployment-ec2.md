# Deploy go-ec2 with DuckDNS and HTTPS

Give your running Go application a URL such as:

```text
https://go-ec2-thuta.duckdns.org
```

This guide uses your Ubuntu EC2 instance, Docker container `go-ec2-container`, and application port `4040`. Run terminal commands on EC2 unless a step says otherwise. The example domain must be registered first; if it is unavailable, choose another name and replace it throughout this guide.

## How requests reach your Go app

```text
Browser → DuckDNS resolves the name to EC2's public IP
        → Caddy on EC2 receives HTTPS on port 443
        → Caddy forwards the request to localhost:4040
        → Docker forwards it to the Go application
```

DNS maps the name to an IP address. Caddy is the reverse proxy that handles HTTPS and forwards requests to your application. [Caddy reverse proxy documentation](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy)

## 1. Verify your application first

In the EC2 terminal:

```bash
docker ps
curl --fail --show-error http://127.0.0.1:4040/
```

The container should be running, publish `4040:4040`, and return your handler's response. The GitHub Actions secret `GO_EC2_PORT` must be `4040` for the current workflow's port mapping.

If the local request fails, inspect the app before configuring DNS:

```bash
docker logs --tail 100 go-ec2-container
```

## 2. Register your free DuckDNS name

1. Open [DuckDNS](https://www.duckdns.org/) and sign in.
2. Add the subdomain `go-ec2-thuta`, if available.
3. In AWS Console, open **EC2 → Instances → your instance** and copy its **Public IPv4 address**.
4. In DuckDNS, enter that address as the domain's IPv4 address and save/update it.

Use the public address. The `172.31.37.153` address shown in your terminal is private and cannot route visitors from the internet. [AWS instance IP addressing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-instance-addressing.html)

DuckDNS provides the free subdomain, so you do not need to buy a domain or create a Route 53 hosted zone for this setup. [DuckDNS](https://www.duckdns.org/)

From your Mac, check the DNS result:

```bash
dig +short A go-ec2-thuta.duckdns.org
```

The result must match the instance's current public IPv4 address. Allow time for DNS caches to refresh before continuing. For this IPv4 setup, leave the DuckDNS IPv6 address empty unless you have configured working public IPv6 connectivity on EC2.

## 3. Open HTTP and HTTPS in the EC2 security group

Open **EC2 → Instances → your instance → Security → attached security group → Edit inbound rules**.

Add these rules:

| Type  | Protocol | Port | Source      |
| ----- | -------- | ---- | ----------- |
| HTTP  | TCP      | 80   | `0.0.0.0/0` |
| HTTPS | TCP      | 443  | `0.0.0.0/0` |

These allow visitors to reach Caddy. Public access to port `4040` is unnecessary for this setup. Keep your SSH access rule suitable for your connection method; for direct SSH from your Mac, restrict it to your public IP. [AWS security group examples](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html)

If you already enabled Ubuntu's UFW firewall, check its status:

```bash
sudo ufw status
```

If it is active, allow web traffic:

```bash
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
```

## 4. Install Caddy on EC2

Use Ubuntu's packaged Caddy:

```bash
sudo apt-get update
sudo apt-get install -y caddy
caddy version
```

The package is available in Ubuntu 26.04's Universe repository. [Ubuntu Caddy package](https://packages.ubuntu.com/en/resolute/caddy)

If APT reports that it cannot locate the package, enable Universe and retry:

```bash
sudo apt-get install -y software-properties-common
sudo add-apt-repository -y universe
sudo apt-get update
sudo apt-get install -y caddy
```

## 5. Configure the domain and reverse proxy

Open the Caddy configuration:

```bash
sudo nano /etc/caddy/Caddyfile
```

On this learning instance, replace the default contents with:

```caddyfile
go-ec2-thuta.duckdns.org {
    reverse_proxy 127.0.0.1:4040
}
```

If you chose another DuckDNS name, use it here. Save with **Ctrl+O**, press **Enter**, then exit with **Ctrl+X**.

Validate the configuration, start Caddy, and reload it:

```bash
sudo caddy validate --config /etc/caddy/Caddyfile --adapter caddyfile
sudo systemctl enable --now caddy
sudo systemctl reload caddy
sudo systemctl status caddy --no-pager
```

Only continue past validation if it succeeds. Caddy obtains and renews a public TLS certificate and redirects HTTP to HTTPS when DNS points to this server and ports `80` and `443` are reachable. [Caddy automatic HTTPS](https://caddyserver.com/docs/automatic-https)

## 6. Test your public URL

From your Mac or browser, visit:

```text
https://go-ec2-thuta.duckdns.org/
https://go-ec2-thuta.duckdns.org/items
https://go-ec2-thuta.duckdns.org/randomuser
```

You can also test from your Mac's terminal:

```bash
curl --fail --show-error https://go-ec2-thuta.duckdns.org/
curl --fail --show-error https://go-ec2-thuta.duckdns.org/randomuser
curl -I http://go-ec2-thuta.duckdns.org/
```

Expect your application's response over HTTPS and an HTTP redirect to HTTPS. The `/randomuser` handler also needs its upstream API to respond successfully.

## GitHub Actions deployments

Your existing workflow in [go-ec2.yml](../.github/workflows/go-ec2.yml) deploys a container on port `4040`. Caddy runs separately on the EC2 host, so it can keep forwarding to that port after a deployment. Container replacement can briefly interrupt requests.

Once HTTPS works, you can update the workflow's **Run docker container** command to bind the app only to localhost and restart it after a host reboot:

```bash
docker run -d \
  --restart unless-stopped \
  -p 127.0.0.1:4040:4040 \
  --name go-ec2-container \
  thutasann/go-ec2:latest
```

Apply this in the workflow and deploy through GitHub Actions; do not run it alongside the existing container with the same name. Caddy will still connect to `127.0.0.1:4040`.

## If the URL does not work

| Symptom                                         | Check                                                                                                      |
| ----------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| Domain does not resolve or returns the wrong IP | Confirm the name exists in DuckDNS, update its public IPv4 address, and repeat `dig`.                      |
| Connection times out                            | Check security group ports `80`/`443`, any active host firewall, and the instance's public internet route. |
| HTTPS certificate fails                         | Check DNS A/AAAA records, port reachability, and Caddy logs.                                               |
| Caddy returns `502 Bad Gateway`                 | Run the local `curl` check and inspect Docker logs; the Go app must listen on port `4040`.                 |
| Caddy fails to start                            | Validate the Caddyfile and check whether another service already uses ports `80` or `443`.                 |

Useful EC2 diagnostics:

```bash
sudo journalctl -u caddy --no-pager -n 100
sudo ss -ltnp
docker ps -a
docker logs --tail 100 go-ec2-container
```

## After stopping and starting EC2

An automatically assigned public IPv4 address changes when you stop and start the instance. Update DuckDNS with the new address, then check DNS again. An Elastic IP can keep the address stable, but AWS charges for public IPv4 addresses, including Elastic IPs. EC2, storage, and traffic costs still depend on your account's plan and credits; the free DuckDNS name does not make AWS usage free. [AWS instance IP addressing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-instance-addressing.html)

This document is a setup guide. Creating the file does not register the domain or configure your EC2 instance.
