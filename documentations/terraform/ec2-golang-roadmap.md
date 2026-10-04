# Learn EC2 by Deploying a Go App

Learn enough Amazon EC2 to run your own Go app, then use Terraform to create the infrastructure again. Start with one Linux server and a tiny HTTP API so you can keep coding while learning cloud basics.

Suggested order: **EC2 basics → local Go API → manual deployment → Terraform**. The sessions below are a suggested pace; repeat a session when you need more practice.

Resource links were checked on October 4, 2026.

## YouTube tutorials to watch first

Start with the EC2 overview, then watch the relevant sections of the hands-on tutorial.

| Order                | Free video                                                                                                 | What to watch                                                                                                                                                                                     |
| -------------------- | ---------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1                    | [Tiny Technical Tutorials — Amazon AWS EC2 Basics](https://www.youtube.com/watch?v=eaicwmnSdCs)            | Watch the opening overview and explanations of AMIs, instance types, storage, security groups, and termination. The demonstration uses Windows and Remote Desktop; use Linux for your Go project. |
| 2                    | [Kevin Stratvert channel — AWS Tutorial for Beginners](https://www.youtube.com/watch?v=Nzv-tzU-UAw)        | Watch the console and EC2 sections, then the connection and cost-saving sections. The tutorial is hosted by Mike Fisher.                                                                          |
| Optional Go practice | [freeCodeCamp — Learn Go Programming by Building 11 Projects](https://www.youtube.com/watch?v=jFfo23yIWac) | Pick a small project for coding practice. You do not need to complete the full course before deploying your first Go app.                                                                         |

Useful timestamps for the Kevin Stratvert video:

- **02:55** — AWS console overview.
- **04:14** — What EC2 is.
- **04:47–07:59** — Launching an EC2 instance.
- **15:39–19:48** — Connecting and working on the instance. The commands relate to the video's larger website project; adapt the idea to your Go app.
- **20:46** — Cost-saving tips.

For this first exercise, focus on one EC2 server. The video's S3 and RDS setup belongs to its larger project and is optional for your Go API.

Older videos can show different console layouts and Free Tier terms. Use the current [AWS EC2 getting started guide](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EC2_GetStarted.html) for the actual launch steps.

## EC2 concepts you need now

| Concept           | Meaning for your project                                                                   |
| ----------------- | ------------------------------------------------------------------------------------------ |
| EC2 instance      | The virtual machine that runs your Go app.                                                 |
| AMI               | The machine image used to launch it, including its operating system.                       |
| Instance type     | The machine's CPU, memory, and other capabilities.                                         |
| Region            | The AWS location where you create the instance.                                            |
| VPC and subnet    | The network and network segment containing your instance.                                  |
| Security group    | Rules controlling traffic to and from the instance.                                        |
| Key pair and SSH  | Credentials and a connection method for accessing a Linux server.                          |
| EBS volume        | Persistent block storage, commonly used for the instance's root disk.                      |
| Public IP address | An address you can use to reach the instance over the internet when networking permits it. |

Sources: [AWS EC2 getting started](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EC2_GetStarted.html) and [EC2 security groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-security-groups.html).

## Session 1 Learn EC2 and connect to Linux

Allow roughly 60–90 minutes.

1. Watch the recommended EC2 video sections.
2. Check your AWS account plan and credit balance before launching a server. Your earlier billing screenshot showed zero available credits, so the signup credit was not confirmed there.
3. Follow the official tutorial to launch one small Linux instance. Use Amazon Linux or Ubuntu, and check the displayed eligibility and estimated costs for your account.
4. Connect using the instructions shown under the instance's **Connect** tab.
5. Practice `pwd`, `ls`, `whoami`, `uname -m`, and `df -h`.

Checkpoint: explain your chosen AMI, instance type, region, connection method, and security group.

For SSH from your laptop, restrict port 22 to your public IP. Browser-based EC2 Instance Connect has its own network requirements; follow the [official connection instructions](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-connect-methods.html) for the method you choose.

## Session 2 Build a small Go API locally

Allow roughly 60–90 minutes. Use your existing Go knowledge to write a small app with these suggested endpoints:

| Endpoint      | Suggested behavior                             |
| ------------- | ---------------------------------------------- |
| `GET /hello`  | Return a greeting as JSON.                     |
| `GET /health` | Return HTTP 200 with a simple health response. |

Use `net/http` for a minimal app, or follow the [official Go and Gin API tutorial](https://go.dev/doc/tutorial/web-service-gin) if you prefer a guided exercise. Test the app locally with `curl`, then add request logging and a configurable listening port.

Checkpoint: you can run the app locally and explain how a request reaches its handler.

## Session 3 Run your Go app on EC2

Allow roughly 60–90 minutes.

1. Build your app for Linux, matching the instance's CPU architecture. A macOS executable cannot run directly on a Linux instance. Check `uname -m` on the server: `x86_64` corresponds to Go's `amd64`, while `aarch64` corresponds to `arm64`.
2. Copy the executable to your server using `scp` or another transfer method you understand.
3. Run it as your normal Linux user on port 8080.
4. Verify the endpoint from inside the instance first.
5. For direct access from your laptop, have the app listen on `:8080` rather than only `localhost:8080`, and allow inbound TCP 8080 from your public IP in the security group.
6. Call the instance's public IP on port 8080 from your laptop and check the response and logs.

Checkpoint: a request from your laptop reaches the Go app running on EC2. Keep this initial exercise to test data over HTTP.

For build settings, see the [Go command environment documentation](https://pkg.go.dev/cmd/go#hdr-Environment_variables). If you adapt the official Gin tutorial, change its localhost-only listening address for direct network access.

## Session 4 Practice an update and clean up

Change the greeting, rebuild, upload the new executable, and restart the app. After that works, learn how a `systemd` service keeps the app running independently of your SSH session.

When you finish the lab, terminate the instance and check for leftover volumes, snapshots, and any allocated Elastic IP addresses. Stopping an instance ends instance compute charges while it is stopped, but retained storage can still incur costs. Termination is permanent, and storage cleanup depends on the volume settings. [AWS instance lifecycle](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)

Checkpoint: you understand how to update your app and remove the resources created for the exercise.

## Then begin Terraform

Once you can deploy manually, follow the [HashiCorp AWS getting started tutorials](https://developer.hashicorp.com/terraform/tutorials/aws-get-started) to create a fresh instance with Terraform.

Recreate the instance and security group you now understand, practice `init`, `plan`, `apply`, and `destroy`, and add outputs for the instance address. Then continue with the [Terraform certification roadmap](README.md).

Your application code defines what the server does. Terraform configuration defines the infrastructure that hosts it. Keep improving your Go app as you learn each infrastructure concept.
