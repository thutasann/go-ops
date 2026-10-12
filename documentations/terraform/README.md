# Practical Go, AWS, and Terraform Learning Roadmap

Learn AWS by deploying a Go app, then rebuild the infrastructure with Terraform. Learn the theory while building, updating, and debugging your project. After you can explain and recreate it independently, use the certification sections to prepare for Terraform Associate (004).

Suggested order: **connect to EC2 → local Go API → deploy and update → Docker and HTTPS → Terraform → certification preparation**. Start with a small working application and add one new concept at a time. Plan for about one hour a day; repeat a session when you need more practice.

## How to learn while building

Use YouTube with your terminal open. Watch 5–10 minutes, pause, and reproduce the steps yourself. After a step works, explain why it was needed, change one setting, and check the result. Rebuild the exercise later without the video.

For a one-hour session, aim for **10 minutes watching, 40 minutes building and debugging, and 10 minutes explaining what you learned**. These are suggested proportions, not a requirement to finish a video or lab in one hour.

Measure progress by working results and your ability to explain them. A completed lab, a useful troubleshooting note, or a project you can rebuild independently is a stronger checkpoint than the number of videos watched.

Keep short notes for each exercise:

- What I built and how I verified it.
- What the important commands and settings do.
- What failed, why it failed, and how I fixed it.
- One change I made without copying the instructor.

## Recommended YouTube tutorials

Start with the practical EC2 lesson, then follow the Terraform project after you can deploy manually. Use overview videos when you need an explanation of a concept encountered during the lab.

| Order                   | Free video or playlist                                                                                                 | Practical goal                                                                                                                        |
| ----------------------- | ---------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| 1                       | [Abhishek Veeramalla — AWS Zero to Hero](https://www.youtube.com/playlist?list=PLdpzxOOAlwvLNOxX0RfndiYSt1Le9azze)     | Start with Day 3: launch EC2, connect with SSH, and deploy an application. Use a small Go API for your exercise.                      |
| After manual deployment | [freeCodeCamp — Learn Terraform and AWS by Building a Dev Environment](https://www.youtube.com/watch?v=iRaai1IBlB0)    | Build a development environment with networking, security groups, and EC2 using Terraform. Follow after your manual deployment works. |
| Optional Go practice    | [freeCodeCamp — Learn Go Programming by Building 11 Projects](https://www.youtube.com/watch?v=jFfo23yIWac)             | Pick a small project for coding practice; you do not need to finish the full course first.                                            |
| Optional explanation    | [TechWorld with Nana — Terraform Explained in 15 Minutes](https://www.youtube.com/watch?v=l5k1ai_GBDE)                 | Use for context on infrastructure as code and Terraform commands during the project.                                                  |
| Optional second lab     | [freeCodeCamp — Terraform Course: Automate Your AWS Cloud Infrastructure](https://www.youtube.com/watch?v=SLB_c_ayRMo) | Practice another guided infrastructure build after finishing your first project.                                                      |
| Exam preparation        | [freeCodeCamp — Terraform Associate Certification Course (003)](https://www.youtube.com/watch?v=SPcwo0Gq9T8)           | Follow after practical work, then cover the 004 updates below.                                                                        |

The AWS playlist includes project code in its [supporting repository](https://github.com/iam-veeramalla/aws-devops-zero-to-hero). The Terraform course is a roughly two-hour project; allow extra time to pause and practice. [Course details](https://www.freecodecamp.org/news/learn-terraform-and-aws-by-building-a-dev-environment/)

The AWS playlist and development environment course were checked on October 12, 2026. Use your selected Region, `ap-southeast-2`, even when a video uses another Region. Use AWS CLI profile `thutasann` explicitly in commands. Your AWS experience manages human access, so a tutorial's IAM user setup may differ from your project.

### Optional EC2 explanations

The following resources were checked on October 4, 2026. Watch only the sections that help with your current exercise.

| Order | Free video                                                                                          | What to watch                                                                                                                                                                                     |
| ----- | --------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1     | [Tiny Technical Tutorials — Amazon AWS EC2 Basics](https://www.youtube.com/watch?v=eaicwmnSdCs)     | Watch the opening overview and explanations of AMIs, instance types, storage, security groups, and termination. The demonstration uses Windows and Remote Desktop; use Linux for your Go project. |
| 2     | [Kevin Stratvert channel — AWS Tutorial for Beginners](https://www.youtube.com/watch?v=Nzv-tzU-UAw) | Watch the console and EC2 sections, then the connection and cost-saving sections. The tutorial is hosted by Mike Fisher.                                                                          |

Useful timestamps for the Kevin Stratvert video:

- **02:55** — AWS console overview.
- **04:14** — What EC2 is.
- **04:47–07:59** — Launching an EC2 instance.
- **15:39–19:48** — Connecting and working on the instance. The commands relate to the video's larger website project; adapt the idea to your Go app.
- **20:46** — Cost-saving tips.

For this first exercise, focus on one EC2 server. The video's S3 and RDS setup belongs to its larger project and is optional for your Go API.

Older videos can show different console layouts and Free Tier terms. Use the current [AWS EC2 getting started guide](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EC2_GetStarted.html) for the actual launch steps.

## Project milestones and the theory behind them

| Build or change                        | Learn while doing it                                         | Evidence it works                                             |
| -------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------- |
| Run a Go API locally                   | HTTP requests, handlers, and listening ports                 | `curl` returns the expected response.                         |
| Connect to an EC2 server               | AMIs, CPU architecture, SSH, and access rules                | You can run commands on the Linux server.                     |
| Deploy the Go executable               | Linux builds, public versus private IPs, and security groups | A request from your laptop reaches the app.                   |
| Update the app and restart it          | Processes, logs, and service management                      | The updated response appears and logs explain the request.    |
| Package the app and add HTTPS          | Docker images, port mappings, DNS, and reverse proxies       | Your application responds through its HTTPS URL.              |
| Recreate a separate lab with Terraform | Providers, resources, dependencies, plans, and state         | You can explain the plan and rebuild the lab without a video. |

Use the repository's [Go app](../../go-ec2/README.md), [Docker installation guide](../../go-ec2/docker-install-ec2.md), and [DuckDNS and HTTPS guide](../../go-ec2/dns-deployment-ec2.md) when you reach those milestones. The HTTPS guide uses port `4040`; the minimal manual exercise below uses `8080`.

## EC2 concepts to look up during the lab

Use this table as a reference when a term appears in the exercise; you do not need to memorize it before starting.

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
2. Check your project plan and spend status in **AWS Settings → Billing** before creating additional resources. Confirm the selected Region in **AWS Settings → View all projects → Overview → Additional Info → Region**.
3. Use your existing learning instance if suitable, or follow the official tutorial to launch one small Linux instance in `ap-southeast-2`. Use Amazon Linux or Ubuntu, and check the displayed eligibility and estimated costs for your project.
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

When you finish the lab, decide whether to keep the instance for your next session or clean up resources created for this exercise. Stopping an instance ends instance compute charges while it is stopped, but retained storage can still incur costs. If you choose termination, it is permanent; check volume deletion settings and any remaining lab volumes, snapshots, or allocated Elastic IP addresses. [AWS instance lifecycle](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)

Checkpoint: you understand how to update your app and remove the resources created for the exercise.

## Rebuild the infrastructure with Terraform

Once you can deploy manually, follow the [HashiCorp AWS getting started tutorials](https://developer.hashicorp.com/terraform/tutorials/aws-get-started) to create a fresh instance with Terraform.

Follow the freeCodeCamp development environment project above and recreate the instance and security group you now understand. Practice `init`, `plan`, `apply`, and `destroy`, and add outputs for the instance address. Use a separate Terraform lab rather than assuming Terraform already manages your manually created server. Read each plan before applying it; use `destroy` only for lab resources you intend to remove. Then continue with the certification preparation below.

Your application code defines what the server does. Terraform configuration defines the infrastructure that hosts it. Keep improving your Go app as you learn each infrastructure concept.

## Terraform certification preparation

Use this stage after completing the deployment and Terraform labs. The suggested six-week plan combines practical exercises with the official exam objectives.

Exam information was checked on October 4, 2026.

The certification video in the tutorial table targets **003**. Supplement it with the current official resources before taking **004**. For older demonstrations, check configuration and provider behavior against the current documentation when your results differ.

### Certification details

The current certification is **HashiCorp Certified: Terraform Associate (004)**.

| Detail                   | Current information                                                              |
| ------------------------ | -------------------------------------------------------------------------------- |
| Terraform version tested | 1.12                                                                             |
| Duration                 | 60 minutes                                                                       |
| Delivery                 | Online proctored                                                                 |
| Format                   | Multiple choice, including true/false and multiple-answer questions              |
| Price                    | US$70.50 plus applicable taxes and fees; a free retake is not included           |
| Credential validity      | 2 years                                                                          |
| Prerequisites            | Basic terminal skills and an understanding of on-premises and cloud architecture |

Sources: [Official exam details](https://docs.hashicorp.com/certifications/terraform-associate) and [official sample question formats](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-questions-004).

### Updates to study for exam 004

HashiCorp lists updates covering these topics:

- `depends_on` and the `create_before_destroy` lifecycle rule.
- Configuration validation using custom conditions.
- Ephemeral values and write-only arguments.
- Organizing and using HCP Terraform workspaces and projects.

The exam tests Terraform 1.12 and includes HCP Terraform content. Use the [official 004 exam content list](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-review-004) to check every objective and close gaps in older courses.

### Suggested six week study plan

This schedule is a recommendation built around the [official 004 learning path](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-study-004). The duration is a planning estimate, not a guarantee of exam readiness.

| Week                               | Learn                                                                                                                          | Practice checkpoint                                                                                            |
| ---------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------- |
| 1 — Foundations                    | Infrastructure as code, basic cloud architecture, providers, resources, HCL syntax                                             | Install Terraform and create, modify, and destroy your first resource.                                         |
| 2 — Core workflow                  | `init`, `fmt`, `validate`, `plan`, `apply`, `destroy`; provider versions and the dependency lock file                          | Read a plan and explain every proposed change before applying it.                                              |
| 3 — Configuration                  | Variables, outputs, locals, data sources, collection types, functions, `count`, `for_each`, dependencies                       | Make your configuration reusable through inputs instead of hardcoded values.                                   |
| 4 — Modules and state              | Module inputs and outputs, module versions, local and remote state, locking, drift, imports, state inspection                  | Create a local module, import an existing resource, and resolve deliberate drift in your practice environment. |
| 5 — HCP Terraform and exam updates | Remote runs, VCS integration, workspaces and projects, collaboration, governance, sensitive-data handling, and the 004 updates | Complete the official HCP Terraform tutorials and review every 004 update.                                     |
| 6 — Exam preparation               | Review all eight objective areas and practice explaining command behavior                                                      | Answer practice questions, investigate mistakes, and rebuild a small project without the video.                |

If you are completely new to cloud infrastructure, allow an additional one or two weeks for terminal basics, networking, virtual machines, storage, and identity/access management. HashiCorp's prerequisites include terminal skills and basic cloud/on-premises architecture knowledge. [Exam prerequisites](https://docs.hashicorp.com/certifications/terraform-associate)

### Practice project

For AWS, build a small project containing a **VPC, subnet, security group, and EC2 web server**. Add variables, outputs, a reusable module, and remote state as your studies progress. Keep the project small enough that you can rebuild and explain it independently.

You can also choose Docker for introductory labs. HashiCorp's learning path includes AWS, Azure, Google Cloud, and Docker options. Provider-specific knowledge is not necessary for the exam. [Official learning path](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-study-004) and [exam content list](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-review-004).

### Exam readiness checklist

Before booking, use this suggested checklist to assess your understanding:

- [ ] Explain infrastructure as code and Terraform's role.
- [ ] Explain providers versus modules, and resources versus data sources.
- [ ] Explain what each core Terraform command does.
- [ ] Read a plan and identify creation, updates, replacement, and destruction.
- [ ] Use variables, outputs, data sources, expressions, and reusable modules.
- [ ] Explain state, remote backends, state locking, and drift.
- [ ] Practice importing resources and inspecting state.
- [ ] Explain sensitive-data handling and the relevant 004 updates.
- [ ] Understand HCP Terraform workflows, workspaces, projects, collaboration, and governance.
- [ ] Review every official objective and work through the official sample questions.
- [ ] Rebuild your small practice project and explain your choices without following a video.

Official resources:

- [Terraform Associate 004 learning path](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-study-004)
- [Terraform Associate 004 exam content list](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-review-004)
- [Terraform Associate 004 sample questions](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-questions-004)
- [Exam details and registration](https://docs.hashicorp.com/certifications/terraform-associate)
