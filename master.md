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

**Cloud**\
AWS, Azure

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
C#, VB.NET, Java, Kotlin, Swift, C, C++, Python, VBA, LabView, Go

**Logging**\
FluentBit, Log4net, Serilog

**Miscellaneous**\
Draw.io, Pandoc, WinMerge

**Microsoft**\
.NET, .NET Aspire, ASP.NET, Power Automate, SignalR, Tye, Windows Forms, WPF, Windows Service, WSL, Windows Terminal, winget

**Mobile**\
Android, iOS, Xamarin

**Monitoring**\
Datadog, Dynatrace, Jaeger, New Relic

**Package Manager**\
Chocolatey, npm, Nuget, pdm, pip, winget, yarn

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
ASP.NET (MVC, Razor Pages, Web Forms, Web API), Bootstrap, CSS, DataTables, Express.js, HTML, HTML Agility Pack (HAP), JavaScript, jQuery, Mantine, Next.js, Nginx, Node.js, Nuxt.js, OWIN, Orval, Postman, React, ReactJS, React router, React Bootstrap, REST, Serve, Silverlight, Swagger (Open API), Telerik UI, TypeScript, Vite, Vue.js, WCF, WordPress

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
**Microsoft Certified**: Azure AI Engineer Associate\
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
{{ if .AfterJan2022 }}
**Director of Cloud and DevOps (Cloud Ops Engineer)**\
CloudSaver, Inc. (January 2022 – Present)
{{- if not .Top5 }}

**Cloud**
- Created a process to launch Cloudsaver on client's AWS commercial and GovCloud accounts.
- Automated shutting down and starting resources using AWS Lambda and Event Bridge.
- Setup multiple environment using EKS, MSK, Aurora, Elasticache Redis, OpenSearch, API Gateway, etc.
- Wrote custom OAuth 2.0 authorizer running on Lambda for API Gateway.
- Configured Ingress-Nginx controller with AWS NLB for EKS and Azure Load Balancer for AKS.
- Configured Glue crawler, Athena and Lambda for Operations team to query invoice data.
- Moved EMR Phoenix backend storage from S3 (EMRFS) to HDFS to improve query performance.
- Configured private connection to Elasticsearch cluster with VPC endpoint and private hosted zone.
- Workaround Micronaut OIDC configuration to enable authentication on AKHQ with Okta. 
- Resolve intermittent Aurora RDS database crash.
- Dual-hosted CloudSaver front end app on Azure Static Web App and AWS EKS.
- Setup automatic wildcard subdomain routing for white labeling.
- Configured cross-region cross-account zero-code bridge from AWS Marketplace SNS to CloudSaver app using SNS, SQS, and EventBridge pipes.
- Processed Azure CUR data daily using Azure Functions and Azure Synapse Analytics with Spark pool.
- Remediated issues to meet CIS 1.4.0. for AWS Foundational Technical Review (FTR).
-	Captured application logs in EKS Fargate to Datadog using sidecar instead of Datadog recommended way of using Lambda.
- Capture APM data and application logs in EKS Nodegroup using Datadog daemonset via unix socket.
- Correlated logs and traces in Datadog by formatting application logs as json.
-	Refactored Docker image to run in both Lambda and ECS Fargate.
-	Enabled/disabled CUR obfuscation flow using CloudWatch composite alarm.
- Updated CUR obfuscation process to use S3 event notification, SQS, and EMR serverless.
-	Automated test file creation and upload using Lambda-backed CloudFormation custom resource when launching resources.
-	Switched Transit Gateway to VPC peering to save data processing cost.
-	Submitted dependencies and multiple python files to EMR serverless using both Python virtual environment and PySpark Native Features.
-	Implemented cluster and application autoscaler for WordPress in ECS.
-	Saved $10,000/month by removing redundancy on legacy application.
- Resolved Aurora MySQL issue due to deadlock caused by index creation.
- Configured pseudo blue/green deployment for API deployment in EKS.
{{- end }}
-	Saved $15,000/week by moving CUR ingestion from EKS Fargate to EKS Nodegroup in public subnet.
-	Bypassed API Gateway 10 MB limit using CloudFront with Python (with pyJwt) custom authorizer running on Lambda@Edge.
{{- if not .Top5 }}

**DevOps**
-	Validated Pull Request source and destination branch to prevent unexpected Git merge in Azure DevOps.
-	Automated comments in Azure DevOps Pull Request using service hooks and Azure Functions.
-	Injected private key for custom authorization using jq and Azure DevOps secure files.
-	Created Azure DevOps extension to move work items between swim lanes by using dropdown.
- Automated moving work items on completed Pull Request between Kanban board columns using Azure Functions.
-	Setup Gated build to validate build before Git merge.
- Implemented GitOps to AWS with change set to prevent unintended change on production stacks.
- Optimized build time by building front and backend in parallel.
{{- end }}
-	Created dynamic Kubernetes workers deployment per Kafka topic using Helm, jq, and yq so workers can be scaled independently.
-	Scheduled pipeline that automatically runs C# script (csx), commit, push, and create Pull Request.
{{- if not .Top5 }}

**Programming**
- Created Cloudsaver CLI using Python, setuptools, and Click.
-	Developed Visual Studio Extension to change the environment that local applications point to.
-	Wrote bash and PowerShell scripts to connect AWS accounts and Azure subscriptions to CloudSaver app.
-	Applied schema migration to all tables using idempotent EF script export, S3 event, Lambdas, and SQS.
-	Update Visual Studio Extension to encrypt/decrypt value with DialogWindow, ToolWindow, and WPF.
-	Saved $7,000/week by refactoring .NET 7 application to run on Linux ECS container.
- Debugged running .NET 7 application using Visual Studio remote debugger.
{{- end }}
-	Developed a web page on CloudSaver app to enable self-configure SSO using React and ASP.NET Web API backend.
{{- if not .Top5 }}

**System**
-	Migrated OpenVPN configuration to new servers.
- Resolved OpenVPN issue where config failed validation but was saved and caused the web interface to be unreachable.
-	Setup KEDA to enable pod scaling in Kubernetes (EKS) based on data from CloudWatch.
-	Enabled monitoring of Kubernetes pods by injecting Dynatrace agents into CloudSaver pods.
- Troubleshot HBase/Phoenix error by using hbck to check for table inconsistencies.
- Configured Fargate logging to send application logs in Kubernetes (EKS) to CloudWatch Logs.
- Configured Jaeger to run in Kubernetes (EKS) in ECS Fargate with and without agent sidecar and OpenSearch as storage.
- Configured Amazon Container Insight using AWS Distro for Open Telemetry (ADOT) in EKS.
- Implemented auto memory dump usng Kubernetes preStop lifecycle.
-	Configured application to export telemetry data directly to Jaeger collector removing the need for sidecar.
-	Patched .NET 6 vulnerability that's no longer fixed using .NET 8 docker image and .NET 6 runtime.
-	Bypassed CoreDNS 5s bug issue by routing pod-to-pod traffic through Ingress.
- Fixed issue with Docker Desktop used significant amount of memory.
{{- end -}}
{{- if not .Top5 }}

**Staff**
- Wrote script to simplified memory dump capture and upload to S3.
- Implemented an auto restart pipeline in ADO.
- Migrated OpenSearch indices to a different cluster using replication.
- Implemented cluster autoscaler on EKS Nodegroup.
- Captured application logs in EKS Nodegroup.
- Added container insights for EKS through add-on.
{{- end -}}
{{- if not .Top5 }}

**Various**
-	Assisted with answering technical questions at CloudSaver booth in AWS re:Invent 2022.
-	Documented various architecture and data flow diagrams using draw.io.
-	Assisted in obtaining SOC 2, ISO 27001, and GDPR compliance.
{{- end -}}
{{- end -}}
{{ if .AfterMay2021 }}

**Software Engineer**\
CloudSaver, Inc. (May 2021 – December 2021)
{{- if not .Top5 }}

**Auth**
-	Setup SAML 2.0 SSO for clients who use Okta, PingFederate, and Azure in Auth0.
-	Fulfilled client security requirements by enabling Single Logout (SLO) between Auth0 and PingFederate.
{{- end }}
-	Updated from built-in authentication to OpenID and OAuth 2.0 using Auth0. 
{{- if not .Top5 }}

**Big Data**
-	Significantly improved ingestion rate from Apache Spark to Apache Phoenix using Phoenix Spark connector.
{{- end }}
-	Wrote C# Apache Spark application running on EMR to reduce Cost and Usage (CUR) data.
-	Launched Apache Phoenix and Apache HBase on EMR to serve invoice data to web application.
{{- if not .Top5 }}

**Cloud**
-	Significantly reduced timeout by migrating database from Amazon Aurora Serverless to Amazon Aurora RDS cluster.
-	Setup WordPress using Bitnami Docker image hosted in ECS Fargate with EFS storage and Amazon Aurora serverless.
-	Launched and configured OpenVPN Access Server.
-	Moved WordPress from ECS Fargate to ECS EC2 to enable faster throughput by connecting it to NFS server backed by EBS.
-	Moved Learning team site to Amazon Lightsail due to storage limit in Azure Static Web App.
- Initialized ASP.NET session database in AWS RDS for SQL Server (web edition) to improve compatibility.
- Resolved issue with disabled AWS SES due to high email bounce rate.
- Re-configured AWS VPC network from VPC peerings to hub and spoke model using Transit Gateway.
- Setup Content Delivery Network (CDN) using Amazon CloudFront and S3.
- Created static website in Amazon S3 and Amazon CloudFront for Learning team.
- Automated copy and bootstrap of Aurora MySQL database to lower environment using EventBridge and Lambda.
- Ran Hangfire as ECS Service and added capability to launch ECS tasks.
- Adopted IaC by launching and importing resources using CloudFormation templates.
{{- end }}
-	Containerized and ran batch applications on ECS Fargate to improve scalability and reduce cost.
{{- if not .Top5 }}

**DevOps**
-	Containerized batch applications on deployment to Amazon ECR (Elastic Container Registry) using Docker.
-	Created Azure Artifacts feed to host private Nuget packages.
-	Utilized AWS CodeDeploy hooks to update running applications on EC2 by renaming files.
-	Installed RDS Combined CA bundle to enable SSL to Aurora MySQL on various resources (Docker, EC2, etc.).
-	Migrated source/version control technology from TFVC to Git.
{{- end -}}
{{- if not .Top5 }}

**Miscellaneous**
- Initiated and led weekly team meeting.
- Initiated and led bi-weekly architecture meeting.
- Setup and maintained Microsoft SharePoint for department information sharing.
-	Implemented Datadog APM for CloudSaver applications on EC2.
-	Replaced Datadog with Dynatrace for CloudSaver applications on EC2 due to APM for .NET was not ready.
{{- end -}}
{{- if not .Top5 }}

**Programming**
-	Configured CloudSaver web application to use in proc session state to fix issue with incompatibility with Aurora MySQL.
-	Upgraded batch applications from .NET 4.6.2 to .NET 5, containerized and run in ECS Fargate.
-	Upgraded shared library from .NET 4.6.2 to .NET Standard 2.0.
-	Created a simple web application using Node.js and VS Code to host temporary learning courses.
-	Used Handlebar.NET to map database bootstrap script with data from AWS secrets manager.
-	Decoupled shared code from batch applications by moving it to a Nuget package.
-	Resolved issue with TLS 1.2 between newer Linux kernel and MySQL by re-enabling older cipher suites.
{{- end }}
-	Wrote EC2 rightsizing scripts requested by client in Bash and PowerShell.
{{- end -}}
{{ if .AfterAug2019 }}

**Consultant**\
CloudSaver, Inc. (August 2019 – May 2021)
{{- if not .Top5 }}

**Auth**
- Built integration with Okta, Ping Identity, OneLogin to enable SAML 2.0 SSO.
- Moved custom Sustainsys SAML 2.0 logic from HttpModule to OWIN Middleware.
- Setup Microsoft Active Directory.
- Enabled Cloudsaver app to authenticate against Microsoft AD.
{{- end }}
{{- if not .Top5 }}

**Cloud**
- Programmatically installed and configured CloudWatch agent using AWS Systems Manager (SSM).
- Used AWS API Gateway and Lambda to provide endpoints to store metadata for CloudSaver bundling and compression engine.
- Automated Amazon EC2 virtualization type conversion from PV to HVM.
- Installed ENA module and NVMe driver programmatically via AWS System Manager (SSM) for Zen to Nitro conversion.
- Saved on SSL certificate cost by using AWS Certificate Manager (ACM).
-	Designed and performed Amazon S3 Bundling PoC for client.
-	Automated creation and sharing of AMI Golden Image with partner.
- Launched NAT Gateway to workaround Elastic IP (EIP) limit.
- Automated benchmarking of EC2 using Aida 64 and Geekbench on custom AMI.
-	Configured Site-to-Site VPN to client environment.
- Setup automatic security patching using AWS SSM Patch Manager
- Setup automatic agents update using AWS SSM State Manager.
{{- end }}
- Implemented EBS Safety Net feature using Amazon CloudWatch alarms and Amazon SNS.
-	Automated benchmarking of EC2 using Aida 64 and Geekbench with custom AMI.
{{- if not .Top5 }}

**Database**
-	Created SQL Server stored procedure with Table-Valued Parameters to generate complex reports.
-	Improved EF batch insert performance by utilizing SqlBulkCopy.
-	Enabled consistent database schema migrations using EF 6 (Entity Framework) Migrations.
- Moved Hangfire database from SQL Server to Amazon Aurora Serverless.
- Moved critical data from SQL Server to Amazon Aurora Serverless.
-	Set up Amazon DocumentDB as a storage to store status and configuration for batch applications.
{{- end }}
-	Implemented Elasticsearch to store and improve query speed of AWS CUR/Invoice data.
{{- if not .Top5 }}

**DevOps**
-	Implemented CI/CD pipelines using Azure DevOps and AWS CodeDeploy with deployment to EC2 and Lambda.
- Designed and applied Azure DevOps building and deployment steps.
- Troubleshot phantom TFS workspace.
- Setup Kanban board, design and document development process in Azure DevOps.
{{- end -}}
{{- if not .Top5 }}

**Graphic**
- Added emboss effect on CloudSaver logo using Adobe Illustrator.
- Edited promotional video using After Effect and Quick Time.
- Assisted in creating and editing of CloudSaver assessment instruction video using Movavi.
- Blurred sensitive information in CloudSaver instruction videos using Adobe Premier Element.
{{- end -}}
{{- if not .Top5 }}

**System**
- Wrote code and scripts to automatically migrate EC2 in Xen hypervisor to Nitro-based hypervisor.
- Automated EBS cost saving by writing scripts to setup and move data to smaller EBS volumes.
- Map disk drives and partitions to EBS volumes using PowerShell and Shell script.
- Created mirror (Raid 1) on EBS using PowerShell and Diskpart.
- Created script to create, attach, initialize, part and format EBS volumes.
- Automated benchmarking by programmatically run AIDA64 on various EC2 using AWS Systems Manager (SSM).
- Hosted private Nuget server in EC2.
- Developed feature to bundle files and extract a single file by querying the byte range in S3.
- Wrote script to replace device name with UUID in fstab to migrate Amazon EC2 to Nitro-based instance type.
- Installed and renewed SSL certificates in IIS.
- Disabled deprecated SSL and TLS versions/protocols on Windows Servers via PowerShell.
- Disabled weak encryption protocols/cipher suites on Windows Servers.
- Configured OpenVPN access server to allow VPN connection to test account VPC.
- Wrote script to automatically fix Linux kernel update failure.
- Troubleshot URL Rewrite on IIS.
{{- end -}}
{{- if not .Top5 }}

**Programming**
- Created .NET generic classes to simplify retrieving multiple metrics from CloudWatch.
- Created .NET extension methods to sort enumerables by attribute and property name using reflection.
-	Implemented drag and drop report filter with level indentation.
- Added custom tooltip on HTML tables using Tooltipster.
- Created custom HTML tables with frozen top, bottom, and first column using CSS.
- Added custom scrollbar using mCustomScrollbar.
- Created various reports in PDF and Excel using TallPDF, EPPlus, and Highcharts.
- Implemented background process queueing and scheduling using Hangfire to decouple front end and back end.
- Created custom .NET attributes to improve code efficiency.
- Worked around template formatting bug in Salesforce CPQ by utilizing knowledge in Apache FOP.
- Developed ASP.NET Core Web API to compress file using LZMA and hosted it in Ubuntu using NGINX.
- Created .NET SDK for CloudSaver compression engine.
- Enabled OWIN on ASP.NET Web Forms.
- Enabled ASP.NET Session on OWIN.
- Solved race condition issue by implementing ServiceLocator pattern.
- Moved some appsettings to environment variable for easier management and security.
- Utilized TopShelf to install Hangfire as Windows Service.
- Improved page response by asynchronously loading batched data using custom queue written in JavaScript.
- Enabled CloudSaver web application to send email using Simple Mail Transfer Protocol (SMTP).
-	Created a Razor page based internal web application to query New Relic API using GraphQL.
-	Automated migration of previous generation Amazon EC2 with Xen hypervisor to newer AWS Nitro hypervisor:
    -	Converted instances from PV (Paravirtual) to HVM (Hardware Virtual Machine)
    -	Installed NVMe driver and ENA (Enhance Network Adapter) module on the instance
    -	Fixed fstab for Linux
- Developed CloudSaver iOS app using Swift and XCode.
{{- end }}
-	Automated moving of data to smaller AWS EBS volumes using shell script, PowerShell and AWS SSM for various OS.
-	Integrated Okta, OneLogin and Ping Identity SAML 2.0 SSO with CloudSaver app using Sustainsys/SAML2 library.
{{- end -}}
{{ if .AfterAug2017 }}

**Web Application Developer**\
WorkView, LLC (July 2017 – August 2019)
{{- if not .Top5 }}


- Enhanced Partner Portal ASP.NET Web Forms application.
- Created windows console app to send out WorkView scheduled reports.
- Added copy button to allow user to copy to clipboard using clipboard.js.
- Created a scheduled report monitoring web application.
- Added logic to extract text from html formatted string using HTML Agility Pack.
- Handled Basecamp3 and Highcharts server rate limitation using custom DelegatingHandler.
- Created various charts using Highcharts.
- Added logic to react when browser tab is visible/invisible using Page Visibility API.
- Added logic to validate URL using Regex.
- Modified images/graphics using Photoshop.
- Automated file copying after code compilation/build using MSBuild.,        
- Troubleshot permissions in SQL Server.
- Managed projects and versions in VSTS.
- Explored jReport as part of reporting tools due diligence.
- Implemented custom logger in ASP.NET Web API applications using log4net.
- Wrote JavaScript query string parser.
- Troubleshot partner company's issue using cURL.
- Developed app to process scheduled reports.
- Moved application error handler to a central location.
- Installed Ubiquity ToughSwitch to enable multiple VoIP phones.
{{- end}}
- Developed and released WorkView mobile app for iOS and Android with ASP.NET Web API backend.
- Generated scheduled reports in PDF (Tallcomponents TallPDF.NET) and in Excel with embedded macro (EPPlus).
- Enabled import of projects and to-dos from Basecamp3 to WorkView.
- Created Gantt chart of projects and tasks programmatically using DayPilot Gantt.
- Added feature to manage meetings and appointments using Exchange Web Services (EWS).
{{- end -}}
{{ if .AfterAug2012 }}

**GUI Programmer**\
Data Center Inc. (August 2012 – August 2017)
-  Developed multiple ASP.NET Web Forms sub applications, including Vendor Management and DCI University.
-  Enabled File Handler Services to combine PDF documents using dbAutotrack PDF-Writer.NET.
-  Modified format of internal periodical reports generated by Windows Service using PDFlib library.
-  Added scrolling capability to Silverlight document viewer plugin.
-  Enable chat capability for Customer Service department by implementing SightMax chat.
-  Designed new and modified existing icons using JASC Paintshop Pro application
-  Implemented the first drag and drop functionality for ASP.NET grid view in main web application.
-  Added embedded PDF viewer to web pages.
-  Shared educational videos on team meetings.
-  Contributed ideas, potential problems and suggestions in spec meetings.
-  Initiated a section on company's newsletter for employees to share their favorite books.
-  Performed code review on modifications by other programmers.
-  Researched and documented necessary steps to upgrade .NET framework versions.
-  Wrote JavaScript for browser bookmark to enable team members to easily navigate to a page by pageID.
-  Crafted complex SQL query to obtain required data.
-  Added build date time to MSBuild script.
-  Solved problem on importing text files due to incompatible encoding scheme.
-  Created image viewer control in HTML5 to replace Silverlight.
-  Refactored File Handler application logic to enable automatic XML serialization and deserialization.
{{- end -}}
{{ if .AfterJul2011 }}

**Production Tech**\
Envision, Inc. through Labor Finders (July 2011 – August 2012)
-	Developed a custom MRP application using Windows Forms in C#.NET on Visual Studio 2008.
-	Utilized Microsoft Access as data storage for custom MRP application.
-	Synchronized data in Microsoft Access with SQL Server using Windows Service.
- Generated requested reports for supervisors and managers by using Crystal Report.
- Deployed custom MRP application using ClickOnce technology.
- Created dynamic charts by utilizing MS (.NET) Chart Control.
- Create a small tool using WPF to keep track of manufactured products.
{{- end -}}
{{ if .AfterAug2008 }}

**Graduate Research Assistant**\
Finance, Real Estate & Decision Sciences Department at WSU (August 2008 – May 2011)
-	Queried real estate data from RETS compliance servers from STATA via ODBC connection (ezRETS driver).
-	Developed statistical analysis programs using STATA to analyze periodic real estate data.
-	Automated reports generation in Microsoft Excel by writing macros in VBA.
- Generated monthly reports to be delivered to 24 Kansas area real estate boards.
- Shortened reports creation time from 1 week to 3 days through automated process.
- Found and reported bug on RETS ODBC driver.
- Integrated real estate data from individual MLS into statewide data.
{{ end }}

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