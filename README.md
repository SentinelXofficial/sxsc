# sxsc -- SentinelX Scanner

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Single-target deep-dive web vulnerability scanner. **46 packages. 79 Go files. 118 YAML blueprints. 15MB binary. 36 attack modules. 19 engines.**

Part of the [SentinelX](https://github.com/SentinelXofficial) ecosystem.

---

## Quick Start

```bash
go install github.com/SentinelXofficial/sxsc/cmd/sxsc@latest

# Basic
sxsc -u "http://testphp.vulnweb.com/listproducts.php?cat=1"

# Lethal mode -- single target, everything on
sxsc -u "http://target.com" --deep --all --prove --delve --vault --rank

# CI/CD
sxsc -u "http://target.com" --all --sarif results.sarif --ci

# Bug bounty pack
sxsc --strike results.json
sxsc --bundle results.json --bundle-out ./submission/

# Diff
sxsc --diff yesterday.json today.json

# Webhook
sxsc -u "http://target.com" --all --hook https://hooks.slack.com/xxx
```

---

## Modules (36)

**Injection:** SQLi (error+blind time+boolean), NoSQLi, Command Injection, SSTI, CRLF, Host Header, Header/Cookie Injection, JSON Injection
**Cross-Site:** XSS, CSRF, CORS (basic+preflight+private network), Open Redirect, Prototype Pollution
**File/Path:** Path Traversal, LFI/RFI+log poisoning, File Upload (9 bypasses), XXE, SSRF (12 probes+timing)
**Auth/Session:** JWT (alg:none, key confusion, weak secret, empty sig), Cookie Audit, IDOR
**Infrastructure:** Security Headers, HTTP Methods, Sensitive Files (47), Directory Brute (190+), GraphQL (5 checks), WAF Detection+Auto-Bypass, Rate Limit Test, TLS Handshake
**Recon:** Subdomain Enum (crt.sh+DNS), Subdomain Takeover (25+ services), WebSocket, JS Endpoints, robots.txt/sitemap
**Advanced:** HTTP Smuggling (CL.TE/TE.CL/TE.TE), Cache Poisoning, Deserialization, Race Condition (TOCTOU), OAuth/SAML Misconfig, gRPC Reflection

---

## Engines (19)

| Engine | Package | Purpose |
|--------|---------|---------|
| Blueprint | `template/` | YAML scans -- Signs, Plucks, 3 fan modes |
| OOB | `oob/` | HTTP+DNS callback for blind vuln confirmation |
| Fuzzer | `fuzzer/` | 12 mutation strategies + grammar generators |
| Auth | `auth/` | Multi-step login with variable extraction |
| Flow | `flow/` | DAG vulnerability chaining |
| Verdict | `verdict/` | Clues + Pickers signal detection |
| Mold | `mold/` | 16-function template variable engine |
| Wire | `wire/` | Raw HTTP byte-level smuggling |
| Strobe | `strobe/` | Adaptive deep-dive pipeline |
| Prove | `prove/` | Auto-demonstrate impact (extract DB, read files, dump IAM) |
| Tally | `tally/` | Composite risk scoring |
| Delve | `delve/` | Auto-escalation (walk IDs, dump tables, extract metadata) |
| Vault | `vault/` | Credential classification |
| Drift | `drift/` | Differential testing (5 Content-Types x 6 methods) |
| Chain | `chain/` | 11 compound attack patterns |
| Pulse | `pulse/` | Session health + auto-renewal |
| Mirror | `mirror/` | Request/response cache + diff |
| Sieve | `sieve/` | Parameter mining from 7 sources |
| Forge | `forge/` | Adaptive payloads based on detected tech stack |

---

## YAML Blueprints (118 in 16 categories)

`cves/` (16) `misconfig/` (15) `exposures/` (12) `technologies/` (16) `panels/` (4) `cloud/` (7) `api/` (6) `cms/` (6) `files/` (6) `devops/` (6) `network/` (6) `frameworks/` (5) `auth/` (4) `backups/` (3) `iot/` (3) `defaults/` (3)

---

## Key Flags

**Target:** `-u`, `-l`, `--cookie`, `-H`, `--proxy`, `--timeout`, `--threads`, `--rate-limit`
**Crawl:** `--crawl`, `--depth`, `--max-pages`, `--scope`, `--out-of-scope`, `--robots`
**Engines:** `--template`, `--template-dir`, `--oob`, `--fuzz`, `--flow`, `--deep`, `--prove`, `--delve`, `--rank`, `--vault`, `--drift`, `--mirror`, `--live`
**Output:** `--html-output`, `--json-output`, `--csv-output`, `--md-output`, `--sarif`, `--bundle`, `--diff`, `--hook`, `--strike`, `--ci`, `--checkpoint`, `--resume`, `--all`

---

## Structure

```
sxsc/
├── cmd/sxsc/             Entry point
├── internal/             Banner, color, updater, version
├── pkg/                  40 packages
│   ├── core/             Foundation: types, HTTP, rate limiter
│   ├── modules/          26 scan modules
│   ├── output/           HTML/JSON/CSV/MD/SARIF reports
│   └── .../              Engines, engines, engines
├── templates/            118 YAML blueprints
├── readme/               10 language docs
└── go.mod
```

---

## Docs

EN | [ID](readme/id.md) | [ZH](readme/zh.md) | [RU](readme/ru.md) | [AR](readme/ar.md) | [ES](readme/es.md) | [JA](readme/ja.md) | [PT](readme/pt.md) | [FR](readme/fr.md) | [HI](readme/hi.md) | [KO](readme/ko.md)

---

## Build

```bash
git clone https://github.com/SentinelXofficial/sxsc
cd sxsc && go mod tidy && go build -o sxsc ./cmd/sxsc
```

---

## Disclaimer

Use only on systems you own or have explicit written authorization. Unauthorized use is illegal.

---

## Author

**WildanDev** -- [@SentinelXofficial](https://github.com/SentinelXofficial)
