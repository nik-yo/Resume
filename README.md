
# Calculation

**Important Metrics**
- Cost savings
- Load handled
- Processing time
- Performance

**NAT Gateway savings** \
$15k/week x 52 weeks/year = $780k/year

**Obfuscation process savings** \
$16.36 on August 2024
EMR on EC2: m5a.xlarge
((0.172 x 3 instances) + 0.043 [EMR]) x 24 x 31 = $415.896
($415.896-$16.36)/$415.896 = 96% savings

**Ticket Velocity** \
2134 tickets from Nov 2022 to May 2024 (team of 2)
Averaging 27 tickets a week or 4 tickets a day

**Marketing requirements** \
Marketing department said the 25,000 emails that they sent crashed the site, so we need to be able to handle that many connections. It was load tested using Locust and autoscaled.

**Azure DevOps self-host** \
Initially 5 parallel jobs, $40 each = $200 a month
Self-host agent, 4 parallel jobs, $15 each = $60 a month
($200-$60)/$200 = 70% savings

**100x increase in performance** \
one pod subscribes to 100 topics instead of only 1 which means processes compete for capacity. By subscribing only to 1 topic, performance increases 100x.

**96% of pod operational cost savings** \
without dynamic scaling, 80 partitions on most active topic x 100 topics = 8000 equivalent capacity. Most topic don't require that much capacity. Dynamic scaling is 80 + (99 x 2) = 278. 278/8000 = 96.5% in saving.

**.NET 7 refactoring savings** \
350k cheaper:
$7k/week x 52 weeks/year = $364k a year

**$6.5 millions in revenue** \
Fannie paid around $4.5m a year in revenue. Merck is around $2m (Merck CUR is about half of Fannie). Fannie Mae has around 100 millions rows (100GB per month), Merck has 50 million rows (50GB per month), 150GB per month x 12 = 1800GB + other companies = 2000GB = 2TB. Data can't be processed by MySql, so need to move to Elasticsearch. Without that migration, there will be no assessment with customers. Without customers, no $6.5 millions.

**Report Generation cost saving**\
jReport quote us around $5k a month, so we code it ourselves.

**SSO cost saving**\
SSO with Fannie is usually 3 x 30 minutes meeting:
1. Initial configuration meeting
2. Troubleshooting meeting
3. Further troubleshooting meeting.
4. Manual work is around 30 minutes to configure Auth0 and identity provider (read documentation, copy ACS URL, upload cert/metadata).
Assuming 1 meeting per week per customer.

**Xen to Nitro time saving**\
- PV to HVM
- NVMe driver
- fstab (Linux)
Manual work measured at 3 hours per instance, automated process in around 30 minutes. 0.5/3 = 83% faster.

**CI/CD benefit**\
- measured at 2 hours manual check and 2 hours of manual deployment and check
- 1 release per week
- CI/CD build runs for 3-5 minutes, deployment 2-5 minutes, 10 minutes max

**ECS Fargate scaling**\
- around 10 batch applications over 100 customers
- scale to around 1000 concurrent execution

**CUR Ingestion time saving**\
100GB of data is ingested in over 24 hours. It is reduced to under 2 hours. 100GB = 100 millions. Spark Phoenix ingestion is around 15-30k rows per second. It was at 24 hours or more. As bad as 28 hours.
(24-2)/24 hours = 92% faster

**Performance analyzer latency**\
Found a multiple network calls which takes 5 minutes and bring it down to under 30 seconds.

**Uptime**\
- 1 minute granularity, 99.99%
- 5 minute, 99.98%
- 1 hour, 99.89%

**API Gateway**\
95%+ in latency improvement:
timeout in API GW is 29s, due to bandwidth, it can't even load, so under 1s rather than 29s is 97% improvement in latency.

**Transit Gateway**\
Transit gateway 6 attachments x $0.05

**EKS Fargate to Nodegroup**\
Saved $21k
- QA: $42.3, $24.7
- Staging: $52.2, $37
- Prod: $49.3, $23.3