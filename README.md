# 🕵️‍♂️ Grepr v2.0.0

Grepr is a blazing-fast, lightweight CLI tool designed for **URL filtering and recon**, built specifically for **bug bounty hunters**, **penetration testers**, and **security researchers**. Inspired by the simplicity of `gf`, Grepr adds targeted intelligence to filter URLs by **file type**, **security templates**, or **super mode automation**.

<p align="center">
  <img src="assets/grepr-demo.gif" alt="Grepr demo" width="600">
</p>

---

## ✨ New in v2.0.0

- 📂 **Template System (GF-style)**: Use predefined security templates like `sqli`, `xss`, `ssrf`, `s3-buckets`, etc.
- 🚀 **Embedded Defaults**: Templates are now embedded in the binary. No more missing file errors!
- 🖥 **Console-First Output**: Results print to console by default (perfect for piping). Use `-o` only when you need to save to a file.
- 🧬 **Dynamic Soora Mode**: Automatically applies all security templates and core filters in one go.
- 🎨 **Modern Banner**: A professional new look for your terminal.

---

## ✨ Features

- 🎯 **Filter by Filetype**: Targeted matching for extensions like `.js`, `.php`, `.aspx`, etc. (Smart enough to handle query params!).
- 🔍 **Regex Filtering**: Filter lines using multiple regex patterns (from a list or file).
- 📋 **Security Templates**: predefined patterns for common vulnerabilities (SQLi, XSS, SSRF, IDOR, LFI, RCE, etc.).
- ⚡️ **Soora Mode**: The ultimate "Deep Scan" mode that runs every filter and template automatically.
- 🪶 **Self-Contained**: Go-embedded templates mean the tool works anywhere without extra setup.
- 📦 **Clean Stats**: Provides line counts and file sizes for all generated outputs.

---

## 📦 Installation

### 1. Using Go Install (Recommended)
This is the fastest way to install Grepr. It will download, compile, and install the binary directly to your `$GOPATH/bin`.

```bash
go install github.com/LaviruDilshan/grepr/v2/cmd/grepr@latest
```

### 2. Build From Source
If you want to contribute or build the binary manually, follow these steps:

```bash
# Clone the repository
git clone https://github.com/LaviruDilshan/grepr.git

# Navigate to the directory
cd grepr

# Build the binary
go build -o grepr main.go

# (Optional) Move to your local bin for global access
sudo mv grepr /usr/local/bin/
```

---

## 🧪 Usage

### Basic Filtering (Console Output)
```bash
grepr -i urls.txt -f js,php
```

### Apply a Security Template
```bash
grepr -i urls.txt -t sqli
```

### List Available Templates
```bash
grepr -l
```

### Save to File
```bash
grepr -i urls.txt -t xss -o results.txt
```

---

## ⚙️ Available Flags

| Flag                   | Description                                               |
| ---------------------- | --------------------------------------------------------- |
| `-i, --input`          | **(Required)** Input file containing URLs                 |
| `-o, --output`         | Output file name (optional, prints to console if omitted) |
| `-f, --filetypes`      | Comma-separated filetypes (e.g., `js,php,aspx`)           |
| `-t, --template`       | Apply a predefined template (e.g., `sqli`, `xss`, `ssrf`) |
| `-l, --list-templates` | List all available security templates                     |
| `-r, --regex-list`     | Regex patterns (comma-separated, e.g., `admin.*,login`)   |
| `--regex-file`         | File path containing regex patterns (one per line)        |
| `-s, --soora`          | Enable **Soora Super Mode** (Applies ALL filters)         |
| `-n, --nobanner`       | Disable the startup banner                                |

---

## 🧠 Soora Super Mode

**Soora** is the ultimate automated engine. When enabled, it performs a comprehensive scan by applying:
1.  **Core Filters**: Extracts all `.js` and `.txt` files.
2.  **Custom Configs**: Applies patterns from `config/extensions.txt` and `config/regex.txt`.
3.  **Security Templates**: Runs **EVERY** available template (`sqli`, `xss`, `idor`, `lfi`, `rce`, `ssrf`, `ssti`, `redirect`, `s3-buckets`, `debug_logic`).

### 🛠️ Usage
```bash
grepr -i all-urls.txt -s
```

### 📁 Output Files
Soora generates individual files for each category and merges them into:
```
📄 Final-Grepr.txt
```

---

## 🔧 Developer Info

* 👨‍💻 **Developer**: Laviru Dilshan  
* 🏢 **Company**: Ovate Security  
* 🌐 **Website**: [lavirudilshan.com](https://lavirudilshan.com)  
* 🐦 **X (Twitter)**: [@LaviruDilshan](https://x.com/LaviruDilshan)  

---

## 🛡️ License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 💬 Feedback

Found a bug or have an idea?
Open an issue or reach out via socials. Contributions are welcome!
