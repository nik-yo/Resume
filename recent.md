# Nikki Yodo

www.nikkiyodo.com | github.com/nik-yo | nik_yodo@yahoo.com | (316) 993-9775

## Technology Experience
**AWS**\
API Gateway, Athena, Aurora, CloudFront, CloudWatch, DocumentDB, EBS, EC2, ECR, ECS, EFS, EKS, EMR, ElastiCache Redis, EventBridge, GuardDuty, Inspector, Lightsail, OpenSearch Service, QuickSight, RDS, Route 53, SES, SQS, SNS, S3, VPC, ALB, ACM, CloudFormation, CodeDeploy, ADOT, Fargate, Glue, IAM, KMS, Lambda, NLB, Private Link, Secrets Manager, STS, Site-to-Site VPN, Step Functions, Systems Manager, Transit Gateway

**Azure**\
App Service, Artifacts, DevOps Services, Functions, Key Vault, Static Web Apps (SWA), Storage Account, Storage Explorer, Synapse Analytics, Bicep

**Auth**\
Auth0, ASP.NET Membership, ASP.NET Identity, IdentityServer, MSAL, OAuth 2.0, One Login, OpenID Connect (OIDC), Okta, Ping Identity, SAML 2.0

**Big Data**\
Apache HBase, Apache Phoenix, Apache Spark, Hue, PySpark, SparkSQL, SQLLine

**Containers**\
Docker, Docker Desktop, Helm, Ingress-Nginx Controller, Kubernetes, KEDA, Minikube

**Database**\
Redis, MongoDB, MySQL, Microsoft SQL Server, SQLite, Squirrel SQL

**Data**\
ADO.NET, AKHQ, Elasticsearch, Entity Framework, GraphQL, JSON, Kafka, LINQ, ODBC, OpenSearch, SQL, XML, XAML, YAML

**IaC**\
CDK, CloudFormation, Bicep, Terraform

**IDE**\
Android Studio, Eclipse, IntelliJ IDEA, Visual Studio, VS Code, Xcode

**Languages**\
C#, VB.NET, Java, Kotlin, Swift, C, C++, Python, VBA, LabView

**Logging**\
FluentBit, Log4net, Serilog

**Miscellaneous**\
Draw.io, WinMerge

**Microsoft**\
.NET, .NET Aspire, ASP.NET, Power Automate, SignalR, Tye, Windows Forms, WPF, Windows Service, WSL, Windows Terminal, winget

**Mobile**\
Android, iOS, Xamarin

**Monitoring**\
Datadog, Dynatrace, Jaeger, New Relic

**Security**\
Bncert, Certbot, Certify the Web, Doppler, Keepass, Let’s Encrypt, Password Agent, Nmap

**Scheduling**\
Hangfire, Quartz.NET

**Scripting**\
PowerShell, PowerShell Core, Bash, Shell Script, jq, yq

**System**\
Forever, Nodemon, Putty, RDP, SSH, WinSCP, WSL

**Version Control**\
Azure DevOps, BitBucket, Git, GitHub, PVCS, Sourcetree, TFVC

**Web**\
ASP.NET (MVC, Razor Pages, Web Forms, Web API), Bootstrap, CSS, DataTables, Express.js, HTML, HTML Agility Pack (HAP), JavaScript, jQuery, Mantine, Next.js, Nginx, Node.js, Nuxt.js, OWIN, Orval, Postman, React, React router, React Bootstrap, REST, Serve, Silverlight, Swagger (Open API), Telerik UI, TypeScript, Vite, Vue.js, WCF, WordPress

## Technical Certifications
**AWS**\
**AWS Certified DevOps Engineer** – Professional\
**AWS Certified Solutions Architect** – Professional\
**AWS Certified Advanced Networking** – Specialty\
**AWS Certified Data Analytics** – Specialty\
**AWS Certified Database** – Specialty\
**AWS Certified Security** – Specialty\
**AWS Certified Developer** – Associate\
**AWS Certified SysOps Administrator** – Associate\
**AWS Certified Solution Architect** – Associate\
**AWS Certified Cloud Practitioner**

**FinOps Foundation**\
FinOps Certified Engineer

**GCP**\
**Google Cloud Certified** – Cloud Digital Leader

**Microsoft**\
**Microsoft Certified**: Azure Solutions Architect Expert\
**Microsoft Certified**: DevOps Engineer Expert\
**Microsoft Certified**: Azure Administrator Associate\
**Microsoft Certified**: Azure Developer Associate\
**Microsoft Certified**: Azure AI Fundamentals\
**Microsoft Certified**: Azure Data Fundamentals\
**Microsoft Certified**: Azure Fundamentals\
**Microsoft Certified Professional**\
**Microsoft Certified Technology Specialist**: .NET Framework 4, Web Applications\
**Microsoft Specialist**: Programming in HTML5 with JavaScript and CSS3\
**Microsoft Technology Associates**: Software Development Fundamental\
**Microsoft Technology Associates**: Database Administration Fundamentals

**Oracle**\
**Oracle Cloud Infrastructure 2023** Certified Foundations Associate\
**Oracle Cloud Infrastructure 2024** Generative AI Certified Professional

## Work Experience

**Director of Cloud and DevOps (Cloud Ops Engineer)**\
CloudSaver, Inc (January 2022 – Present)

**Cloud**
-	Saved $15,000/week by moving CUR ingestion from EKS Fargate to EKS Nodegroup in public subnet.
- Wrote custom OAuth 2.0 authorizer running on Lambda for API Gateway.
- Configured private connection to Elasticsearch cluster with VPC endpoint and private hosted zone.
- Workaround Micronaut OIDC configuration to enable authentication on AKHQ with Okta. 
- Setup automatic wildcard subdomain routing for white labeling.
- Configured zero code bridge between AWS Marketplace SNS topic and CloudSaver app in different AWS regions using SNS, SQS, and EventBridge pipes.
-	Bypassed API Gateway 10 MB limit using CloudFront with custom OAuth 2.0 authorizer written in Python (with pyJwt) running on Lambda@Edge.
-	Captured application logs in EKS Fargate to Datadog using sidecar instead of Datadog recommended way of using Lambda.
-	Refactored Docker image to run in both Lambda and ECS Fargate.
-	Enabled/disabled CUR obfuscation flow using CloudWatch composite alarm.
-	Automated test file creation and upload using Lambda-backed CloudFormation custom resource when launching resources.
-	Switched Transit Gateway to VPC peering to save data processing cost.
-	Enabled submission of dependencies and multiple python files to EMR serverless using both Python virtual environment and PySpark Native Features.
-	Implemented cluster and application autoscaler for WordPress in ECS.
-	Saved $10,000/month by removing redundancy on legacy application.

**DevOps**
- Deployed React front end to both Azure Static Web App and AWS EKS.
-	Validated Pull Request source and destination branch to prevent unexpected Git merge in Azure DevOps.
-	Automated comments in Azure DevOps Pull Request using service hooks and Azure Functions.
-	Injected private key for custom authorization using jq and Azure DevOps secure files.
-	Created Azure DevOps extension to move work items between swim lanes by using dropdown.
- Automated moving work items associated with completed Pull Request between Kanban board columns.
-	Created dynamic Kubernetes workers deployment per Kafka topic using Helm, jq, and yq so workers can be scaled independently.
-	Scheduled pipeline that automatically runs C# script (csx) to pull Azure resource metadata, commit, push, and create Pull Request.
-	Setup Gated build to validate build before Git merge.

**Programming**
-	Developed Visual Studio Extension to change the environment that local applications point to.
-	Wrote bash and PowerShell scripts to connect AWS accounts and Azure subscriptions to CloudSaver app.
-	Developed a web page on CloudSaver app to self-configure SSO using React and ASP.NET Web API backend.
-	Wrote applications to apply schema migration to all tables using idempotent EF script export, S3 event, Lambdas, and SQS.
-	Saved $7,000/week by refactoring .NET 7 application to run on Linux ECS container.
-	Update Visual Studio Extension to encrypt/decrypt value with DialogWindow, ToolWindow, and WPF.

**System**
-	Migrated OpenVPN configuration to new servers.
-	Setup KEDA to enable pod scaling in Kubernetes (EKS) based on data from CloudWatch.
-	Enabled monitoring of Kubernetes pods by injecting Dynatrace agents into CloudSaver pods.
-	Configured application to export telemetry data directly to Jaeger collector removing the need for sidecar.
-	Patched .NET 6 vulnerability that's no longer fixed using .NET 8 docker image and .NET 6 runtime.
-	Bypassed CoreDNS 5s bug issue by routing pod-to-pod traffic through Ingress.
- Configured Datadog APM by launching Datadog agent as sidecar for EKS Fargate and correlated between traces and logs.
- Use Datadog operator to capture application traces in EKS Nodegroup.
- Fixed issue with Docker Desktop used significant amount of memory.

**Various**
-	Assisted with answering technical questions at CloudSaver booth in AWS re:Invent 2022.
-	Documented various architecture and data flow diagrams using draw.io.
-	Assisted in obtaining SOC 2, ISO 27001, and GDPR compliance.


**Software Engineer**\
CloudSaver, Inc (May 2021 – December 2021)

**Auth**
-	Updated from built-in authentication to OpenID and OAuth 2.0 using Auth0. 
-	Setup SAML 2.0 SSO for clients who use Okta, PingFederate, and Azure in Auth0.
-	Fulfilled client security requirements by enabling SLO (Single Log Out) between Auth0 and PingFederate.

**Big Data**
-	Implemented Apache Spark using EMR to reduce invoice/CUR data for web app consumption.
-	Utilized Apache Phoenix and Apache HBase on EMR to serve invoice data to web application. 
-	Significantly improved ingestion rate from Apache Spark to Apache Phoenix using Phoenix Spark connector.

**Cloud**
-	Containerized and ran batch applications on ECS Fargate to improve scalability and reduce cost.
-	Significantly reduced timeout by migrating database from Amazon Aurora Serverless to Amazon Aurora RDS cluster.
-	Setup WordPress using Bitnami Docker image hosted in ECS Fargate with EFS storage and Amazon Aurora serverless.
-	Launched and configured OpenVPN Access Server.
-	Moved WordPress from ECS Fargate to ECS EC2 to enable faster throughput by connecting it to NFS server backed by EBS.
-	Moved Learning team site to Amazon Lightsail due to storage limit in Azure Static Web App.

**DevOps**
-	Migrated source/version control technology from TFVC to Git.
-	Containerized batch applications on deployment to Amazon ECR (Elastic Container Registry) using Docker.
-	Created Azure Artifacts feed to host private Nuget packages.
-	Utilized AWS CodeDeploy hooks to update running applications on EC2 by renaming files.
-	Installed RDS Combined CA bundle to enable SSL to Aurora MySQL on various resources (Docker, EC2, etc.).

**Monitoring**
-	Implemented Datadog APM for CloudSaver applications on EC2.
-	Replaced Datadog with Dynatrace for CloudSaver applications on EC2 due to APM for .NET was not ready.

**Programming**
-	Configured CloudSaver web application to use in proc session state to fix issue with incompatibility with Aurora.
-	Upgraded batch applications from .NET 4.6.2 to .NET 5 (with AWS SDK for .NET Core).
-	Upgraded shared library from .NET 4.6.2 to .NET Standard 2.0.
-	Wrote EC2 rightsizing scripts requested by client in Bash and PowerShell.
-	Created a simple web application using Node.js and VS Code to host temporary Learning courses.
-	Used Handlebar.NET to map database bootstrap script with data from AWS secrets manager.
-	Decoupled shared code from batch applications by moving it to a Nuget package.
-	Resolved issue with TLS 1.2 between newer Linux kernel and MySQL by re-enabling older cipher suites.


**Consultant**\
CloudSaver Inc. (August 2019 – May 2021)

**Cloud**
- Implemented EBS Safety Net feature using Amazon CloudWatch alarms and Amazon SNS.
-	Designed and performed Amazon S3 Bundling PoC for client.
-	Automated creation and sharing of AMI Golden Image with partner.
-	Automated benchmarking of EC2 using Aida 64 and Geekbench with custom AMI.
-	Configured Site-to-Site VPN to client environment.

**Database**
-	Created SQL Server stored procedure with Table-Valued Parameters to generate complex reports.
-	Improved EF batch insert performance by utilizing SqlBulkCopy.
-	Enabled consistent database schema migrations using EF 6 (Entity Framework) Migrations.
-	Implemented Elasticsearch to store and improve query speed of AWS CUR/Invoice data.
-	Set up Amazon DocumentDB as a storage to store status and configuration for batch applications.

**DevOps**
-	Implemented CI/CD pipelines using Azure DevOps and AWS CodeDeploy with deployment to EC2 and Lambda.

**Programming**
-	Implemented drag and drop report filter with level indentation.
-	Automated migration of previous generation Amazon EC2 with Xen hypervisor to newer AWS Nitro hypervisor:
    -	Converted instances from PV (Paravirtual) to HVM (Hardware Virtual Machine)
    -	Installed NVMe driver and ENA (Enhance Network Adapter) module on the instance
    -	Fixed fstab for Linux
-	Created a Razor page based internal web application to query New Relic API using GraphQL.
-	Automated moving of data to smaller AWS EBS volumes using shell script, PowerShell and AWS SSM for various OS.
-	Integrated SAML 2.0 SSO handling using Sustainsys/SAML2 library on CloudSaver web application for Okta, OneLogin and Ping Identity.


**Web Application Developer**\
WorkView, LLC. (August 2017 – August 2019)
-	Developed and released WorkView mobile app for iOS and Android with ASP.NET Web API backend.
-	Generated and scheduled reports in PDF using Tallcomponents TallPDF.NET and in Excel with embedded macro using EPPlus.
-	Added functionality to manage meetings and appointments using Exchange Web Services (EWS).
-	Created Gantt chart of projects and tasks programmatically using DayPilot Gantt.
-	Automated file copying after code compilation/build using MSBuild.





## Education
**Master of Science in Computer Science**\
Wichita State University, Wichita, KS\
May 2011\
GPA: 3.9/4.0

**Bachelor of Science in Aerospace Engineering**\
Wichita State University, Wichita, KS\
December 2007\
GPA: 3.6/4.0

## Activities & Honors
Eta Kappa Nu Honor Society (2010 – Present)\
The National Scholars Honor Society (2007 – Present)\
Golden Key Honor Society (2007 – Present)\
Order of the Engineer (2007 – Present)\
Tau Beta Pi Honor Society (2005 – Present)