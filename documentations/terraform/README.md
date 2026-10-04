# Terraform Associate Certification Roadmap

This guide covers a suggested six-week study plan for HashiCorp Certified: Terraform Associate (004), with free YouTube tutorials, practical exercises, and official exam preparation resources.

Recommended path: **short Terraform introduction → hands-on AWS crash course → certification course → official 004 exam preparation**. Plan for about one hour a day, adjusting based on your cloud experience.

Exam information was checked on October 4, 2026.

If you want to learn EC2 first while continuing to code, start with [Learn EC2 by Deploying a Go App](ec2-golang-roadmap.md), which includes YouTube recommendations and a short deployment learning path.

## Certification details

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

## Recommended YouTube tutorials

Watch these free resources in order, practicing alongside the demonstrations.

| Order | Video or course                                                                                                        | Suggested use                                                                           |
| ----- | ---------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| 1     | [TechWorld with Nana — Terraform Explained in 15 Minutes](https://www.youtube.com/watch?v=l5k1ai_GBDE)                 | Start here for Terraform, infrastructure as code, architecture, and basic commands.     |
| 2     | [freeCodeCamp — Terraform Course: Automate Your AWS Cloud Infrastructure](https://www.youtube.com/watch?v=SLB_c_ayRMo) | Use as your practical crash course. Follow along and build the infrastructure yourself. |
| 3     | [freeCodeCamp — Terraform Associate Certification Course (003)](https://www.youtube.com/watch?v=SPcwo0Gq9T8)           | Use as your longer certification study course, then cover the 004 updates below.        |

The certification video targets **003**. Supplement it with the current official resources before taking **004**. For older demonstrations, check configuration and provider behavior against the current documentation when your results differ.

## Updates to study for exam 004

HashiCorp lists updates covering these topics:

- `depends_on` and the `create_before_destroy` lifecycle rule.
- Configuration validation using custom conditions.
- Ephemeral values and write-only arguments.
- Organizing and using HCP Terraform workspaces and projects.

The exam tests Terraform 1.12 and includes HCP Terraform content. Use the [official 004 exam content list](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-review-004) to check every objective and close gaps in older courses.

## Suggested six week study plan

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

## Practice project

For AWS, build a small project containing a **VPC, subnet, security group, and EC2 web server**. Add variables, outputs, a reusable module, and remote state as your studies progress. Keep the project small enough that you can rebuild and explain it independently.

You can also choose Docker for introductory labs. HashiCorp's learning path includes AWS, Azure, Google Cloud, and Docker options. Provider-specific knowledge is not necessary for the exam. [Official learning path](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-study-004) and [exam content list](https://developer.hashicorp.com/terraform/tutorials/certification-004/associate-review-004).

## Exam readiness checklist

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

## Daily study routine

For each study hour, aim for **20 minutes watching, 30 minutes practicing, and 10 minutes reviewing mistakes**. Start with Nana's introduction, then do the AWS crash course with your terminal open.
