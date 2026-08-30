# Master Resume

This repo holds a single **master resume** (`master.md`) that generates multiple **targeted, role-specific resumes** from one source of truth, instead of maintaining several resume files by hand.

`master.md` is written as a [Go `text/template`](https://pkg.go.dev/text/template). Work history, skills, and certifications are wrapped in conditional blocks so the same document can be filtered two ways:

- **By role** — e.g. `{{ if eq .Role "CloudEngineer" }} ... {{- end }}` includes bullets, skills, and certificates relevant to that target role and omits the rest.
- **By recency** — e.g. `{{ if .AfterJan2022 }} ... {{- end }}` includes older roles only while they still fall within the desired lookback window, so the resume can be trimmed as experience ages.

## Files

| File | Purpose |
| --- | --- |
| `master.md` | The master resume source. Edit this — it's the single source of truth. |
| `generate.go` | Renders `master.md` once per target role and writes each result to `out/`. |
| `out/resume-<Role>.md` | Generated, role-targeted resumes. Not source of truth — regenerated from `master.md`, don't hand-edit. |
| `script.ps1` | Convenience commands: regenerate via `go run generate.go`, plus a Pandoc example for Markdown → PDF conversion. |
| `README.md` | This file, plus the metric calculations below that back up the numbers used in `master.md`. |

## Usage

1. Edit `master.md` — update experience, skills, or metrics, wrapping role- or date-specific content in the appropriate template conditionals.
2. If you add or change a quantified claim (cost savings, performance gains, uptime, etc.), add or update its supporting math in the **Metric Calculations** section below so every number in the resume stays traceable.
3. Regenerate the targeted resumes:
   ```
   go run generate.go
   ```
4. Convert a generated resume to PDF if needed (the VS Code Pandoc extension is recommended over the raw CLI for styling):
   ```
   pandoc -f markdown -t pdf out/resume-CloudEngineer.md -o resume.pdf
   ```

### Adding a new target role

1. In `master.md`, add `{{- if eq .Role "NewRole" }} ... {{- end }}` blocks wherever content should differ for that role.
2. In `generate.go`'s `main()`, add a generation pass for the new role so it gets its own `out/resume-<Role>.md`.

## Metric Calculations

The math behind the quantified claims used throughout `master.md`, kept here so every number in the resume is traceable back to its source.

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
