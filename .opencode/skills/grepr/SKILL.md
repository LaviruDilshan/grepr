---
name: Grepr
description: Fast URL filtering and reconnaissance with grepr CLI. Use when filtering URLs by file type, applying security templates (sqli, xss, ssrf, idor, lfi, rce, ssti, redirect, s3-buckets, debug_logic), running regex patterns against URL lists, or performing comprehensive URL analysis with Soora mode for bug bounty hunting and penetration testing.
---

# Skill: Grepr

Grepr is a blazing-fast CLI tool for URL filtering and reconnaissance, built for bug bounty hunters, penetration testers, and security researchers. It filters URLs by file type, security templates, or custom regex patterns.

## Installation

```bash
# Using Go
go install github.com/LaviruDilshan/grepr/v2/cmd/grepr@latest

# From source
git clone https://github.com/LaviruDilshan/grepr.git
cd grepr
go build -o grepr main.go
sudo mv grepr /usr/local/bin/
```

## Command Reference

```bash
grepr [flags]
```

### Required Flag

| Flag | Description |
|------|-------------|
| `-i, --input <file>` | Input file containing URLs (required) |

### Optional Flags

| Flag | Description |
|------|-------------|
| `-o, --output <file>` | Output file (prints to console if omitted) |
| `-f, --filetypes <types>` | Comma-separated filetypes (e.g., `js,php,aspx`) |
| `-t, --template <name>` | Apply a security template |
| `-l, --list-templates` | List all available templates |
| `-r, --regex-list <patterns>` | Comma-separated regex patterns |
| `--regex-file <file>` | File containing regex patterns (one per line) |
| `-s, --soora` | Enable Soora Super Mode (all filters) |
| `-n, --nobanner` | Disable the startup banner |

## Available Security Templates

List templates with `grepr -l`:

| Template | Purpose |
|----------|---------|
| `sqli` | SQL injection patterns |
| `xss` | Cross-site scripting patterns |
| `ssrf` | Server-side request forgery patterns |
| `idor` | Insecure direct object reference patterns |
| `lfi` | Local file inclusion patterns |
| `rce` | Remote code execution patterns |
| `ssti` | Server-side template injection patterns |
| `redirect` | Open redirect patterns |
| `s3-buckets` | AWS S3 bucket patterns |
| `debug_logic` | Debug and logic flaw patterns |

## Usage Examples

### Filter by File Type

Extract JavaScript and PHP URLs:

```bash
grepr -i urls.txt -f js,php
```

Extract common web file types:

```bash
grepr -i urls.txt -f js,php,aspx,jsp,json
```

### Apply Security Templates

Find potential SQL injection points:

```bash
grepr -i urls.txt -t sqli
```

Find XSS candidates:

```bash
grepr -i urls.txt -t xss
```

Find SSRF candidates:

```bash
grepr -i urls.txt -t ssrf
```

Save results to file:

```bash
grepr -i urls.txt -t sqli -o sqli-candidates.txt
```

### Custom Regex Filtering

Filter with inline regex patterns:

```bash
grepr -i urls.txt -r "admin.*,login,api/v[0-9]+"
```

Filter with regex from file:

```bash
grepr -i urls.txt --regex-file patterns.txt
```

### Soora Super Mode

Run comprehensive analysis with all templates and filters:

```bash
grepr -i urls.txt -s
```

Soora mode performs:
1. Core filters: Extracts `.js` and `.txt` files
2. Custom configs: Applies patterns from `config/extensions.txt` and `config/regex.txt`
3. All security templates: Runs every template automatically

Output: Individual files per category merged into `Final-Grepr.txt`

### Pipeline Integration

Combine with other tools:

```bash
# Filter URLs from subfinder/httpx output
cat subdomains.txt | httpx -silent | grepr -i /dev/stdin -t sqli

# Chain with nuclei
grepr -i urls.txt -t sqli | nuclei -t sqli/ -silent

# Extract JS files for further analysis
grepr -i urls.txt -f js -o js-files.txt
```

Suppress banner for clean piping:

```bash
grepr -i urls.txt -t xss -n | tee xss-results.txt
```

## Workflow: Bug Bounty Reconnaissance

### Step 1: Collect URLs

Gather URLs from various sources (waybackurls, gau, katana, etc.):

```bash
cat all-urls.txt | sort -u > unique-urls.txt
```

### Step 2: Quick Triage with Templates

```bash
# Check for high-value vulnerability patterns
grepr -i unique-urls.txt -t sqli -o sqli.txt
grepr -i unique-urls.txt -t xss -o xss.txt
grepr -i unique-urls.txt -t ssrf -o ssrf.txt
grepr -i unique-urls.txt -t idor -o idor.txt
```

### Step 3: Extract Interesting File Types

```bash
# JavaScript files for secrets/endpoints
grepr -i unique-urls.txt -f js -o js-files.txt

# Config and backup files
grepr -i unique-urls.txt -f json,xml,yml,bak,sql -o config-files.txt
```

### Step 4: Full Analysis with Soora

For comprehensive scanning:

```bash
grepr -i unique-urls.txt -s
# Review Final-Grepr.txt for all findings
```

## Tips

- **Console-first output**: Results print to stdout by default, perfect for piping
- **Smart extension handling**: Correctly handles extensions even with query parameters
- **Embedded templates**: No external files needed; templates are built into the binary
- **Combine filters**: Use multiple `-f` types or chain multiple grepr commands
- **Silent mode**: Use `-n` to disable banner when piping output

## Project Info

- **Repository**: https://github.com/LaviruDilshan/grepr
- **Developer**: Laviru Dilshan
- **Company**: Ovate Security
- **License**: MIT
