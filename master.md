# Nikki Yodo

nik_yodo@yahoo.com | https://www.nikkiyodo.com/ 

https://www.linkedin.com/in/nikkiyodo/ | https://github.com/nik-yo/ 

https://blog.nikkiyodo.com

Overland Park, KS 66210 | (316) 993-9775


## Work Experience

{{ if .AfterJan2025 }}
**Senior Consultant**
Integrity Inspired Solutions (Jan 2025 - Present)

- Build a CLI tool that saves each developer 15 minutes per local app run.
- Rewrite copy process using AI, reducing runtime from timeout to 6 seconds across 200 modules.
- Lead development of a NuGet package with CI/CD pipeline, unit and integration tests, and resiliency testing, delivered within  3 weeks.
- Simplify the local app launch process from 17 steps to 5.
- Troubleshoot TimeZoneInfo discrepancies between Windows (local) and Linux (Kubernetes via Podman and .NET Aspire) environments.
- Integrate Jira with GitLab to surface Merge Request links directly in Jira tickets, improving visibility for developers.
- Develop a randomizer for TPO using Office Script.
- Resolve a memory issue in binary streaming, reducing memory usage by over 90%.
{{- end }}

{{ if .AfterJan2022 }}
**Director of Cloud and DevOps (Hands-on Cloud Ops Engineer)**\
CloudSaver, Inc. (Jan 2022 – December 2024)

- Managed a team of two, completing an average of 27 tickets weekly for a total of 2,134 tickets over 1.5 years.

{{- if or (eq .Role "CloudEngineer") (eq .Role "SolutionArchitect") }}
- Revamped AWS's CUR (Cost and Usage Report) ingestion process using an EKS Nodegroup, saving $780,000 annually in NAT Gateway data processing costs.
- Implemented autoscaling for a WordPress site to support 25,000 concurrent connections per marketing requirements.
- Launched and maintained a large-scale, microservice-based SaaS platform spanning 25 cloud resources across 5 environments, achieving 99.99% uptime since inception.
{{- end }}

{{- if eq .Role "DevOpsEngineer" }}
- Cut costs by 70% by self-hosting Azure DevOps agents in AWS.
- Improved performance by 100x and cut operational costs by 96% by configuring CI/CD pipelines for dynamic deployment and independent scaling of Kubernetes pods per Kafka topic.
{{- end }}

{{- if or (eq .Role "SoftwareEngineer") (eq .Role "MobileDev") }}
- Implemented a CUR obfuscation process in EMR Serverless with Python, achieving a 96% cost reduction compared to EC2-based EMR cluster execution.
- Saved $360,000 annually by refactoring a .NET 7 application to run in a Linux container, leveraging lower ECS Linux rates and improved performance.
- Developed a React and ASP.NET-based web page for customer self-configuration of SAML 2.0 SSO, saving 3 weeks of meetings per customer.
{{- end }}

{{- end }}

{{ if .AfterMay2021 }}
**Software Engineer**\
CloudSaver, Inc. (May 2021 - Dec 2021)

- Led weekly team meetings and bi-weekly architecture meetings, improving Dev and QA collaboration, workload predictability, and reducing technical debt.
- Developed an ETL process from Apache Spark to Apache Phoenix (HBase), reducing CUR ingestion time by 92%, from over 24 hours to under 2 hours.
- Upgraded batch applications to .NET 5 and containerized them for ECS Fargate, enabling scaling to 1,000 concurrent executions.
{{- end }}

{{ if .AfterAug2019 }}
**Consultant**\
CloudSaver, Inc. (Aug 2019 - May 2021)

- Automated the build and deployment process by implementing a CI/CD pipeline, reducing errors and cutting release time from 4 hours to 10 minutes, a 96% improvement.
- Streamlined automation of Xen-to-Nitro EC2 instance conversion, achieving an 83% increase in conversion speed per instance.
- Migrated 2TB and 2 billion rows of data from MySQL to Elasticsearch, reducing query latency and securing two major clients representing $6.5 million in annual revenue.
{{- end }}

{{ if .AfterJul2017 }}
**Web Application Developer**\
WorkView, LLC (Jul 2017 - Aug 2019)
- Delivered Android and iOS apps with an ASP.NET Web API backend in 6 months, completing an iOS project that had been stalled for 9 months.
- Developed in-house applications for report generation, saving the company $60,000 annually.
{{- end }}

## Core Skills
{{- if or (eq .Role "CloudEngineer") (eq .Role "SolutionArchitect") }}
AWS, Azure, CLI, IaC, CloudFormation, Bicep, Docker, Helm, Kubernetes, CI/CD, Git, C#, .NET, Python, VS Code, React, JavaScript, PowerShell, Bash, CDK, MySql, Apache Kafka, Apache Spark, Datadog, Open Telemetry, Golang, Linux, Windows.
{{- end }}

{{- if eq .Role "DevOpsEngineer" }}
Azure DevOps, CI/CD, Git, C#, Python, PowerShell, Bash, Linux, Windows.
{{- end }}

{{- if or (eq .Role "SoftwareEngineer") (eq .Role "MobileDev") }}
C#, .NET, Python, Java, React, JavaScript, VS Code, Kotlin, Swift, PowerShell, Bash, C, C++.
{{- end }}

## Education
**Master of Science in Computer Science**\
Wichita State University, Wichita, KS\
GPA: 3.9/4.0

**Bachelor of Science in Aerospace Engineering**\
Wichita State University, Wichita, KS\
GPA: 3.6/4.0

## Honor Societies
Eta Kappa Nu Honor Society (2010 – Present)\
The National Scholars Honor Society (2007 – Present)\
Golden Key Honor Society (2007 – Present)\
Order of the Engineer (2007 – Present)\
Tau Beta Pi Honor Society (2005 – Present)

## Certificates
{{- if or (eq .Role "CloudEngineer") (eq .Role "SolutionArchitect") }}
**AWS Certified Solutions Architect Professional**\
AWS

**AWS Certified Advanced Networking Specialty**\
AWS

**AWS Certified SysOps Administrator Associate**\
AWS

**AWS Certified Security Specialty**\
AWS

**Microsoft Certified Azure Solutions Architect Expert**\
Microsoft

**Microsoft Certified Azure Administrator Associate**\
Microsoft

**FinOps Certified Engineer**\
FinOps Foundation

**Google Cloud Certified Cloud Digital Leader**\
Google 

**Oracle Cloud Infrastructure 2023 Certified Foundations Associate**\
Oracle

{{- end }}

{{- if eq .Role "DevOpsEngineer" }}
**AWS Certified DevOps Engineer Professional**\
AWS

**Microsoft Certified DevOps Engineer Expert**\
Microsoft

{{- end }}

{{- if or (eq .Role "SoftwareEngineer") (eq .Role "MobileDev") }}
**AWS Certified Developer Associate**\
AWS

**Microsoft Certified Azure Developer Associate**\
Microsoft

**AWS Certified Database Specialty**\
AWS

**AWS Certified AI Practitioner**\
AWS

**Microsoft Certified Azure AI Engineer Associate**\
Microsoft

**Oracle Cloud Infrastructure 2024 Generative AI Certified Professional**\
Oracle

{{- end }}

{{- if .IncludeOtherCert }}
**AWS Certified Solution Architect Associate**\
AWS

**AWS Certified Cloud Practitioner**\
AWS

**AWS Certified Data Analytics Specialty**\
AWS

**Microsoft Certified Azure AI Fundamentals**\
Microsoft

**Microsoft Certified Azure Data Fundamentals**\
Microsoft

**Microsoft Certified Azure Fundamentals**\
Microsoft

**Microsoft Certified Professional**\
Microsoft

**Microsoft Certified Technology Specialist .NET Framework 4, Web Applications**\
Microsoft

**Microsoft Specialist Programming in HTML5 with JavaScript and CSS3**\
Microsoft

**Microsoft Technology Associates Software Development Fundamental**\
Microsoft

**Microsoft Technology Associates Database Administration Fundamentals**\
Microsoft
{{- end }}