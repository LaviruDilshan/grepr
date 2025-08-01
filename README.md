# 🕵️‍♂️ Grepr

Grepr is a blazing-fast, lightweight CLI tool designed for **URL filtering and recon**, built specifically for **bug bounty hunters**, **penetration testers**, and **automation workflows**. Inspired by the simplicity of `grep`, Grepr adds targeted intelligence to filter URLs by **file type**, **regex patterns**, or **super mode automation**.

<p align="center">
  <img src="demo.gif" alt="Grepr demo" width="600">
</p>

---

## ✨ Features

- 🎯 Filter URLs by file types (e.g., `.js`, `.php`, `.aspx`, etc.)
- 🔍 Filter lines using **multiple regex patterns** (from a list or file)
- ⚡️ Soora Mode: Built-in **super filter engine** with curated rules (filetypes, keywords, regex)
- 🪶 Lightweight, fast, and **easy to integrate** into any recon workflow
- 📦 Clean output with **file stats** (line count, file size)
- 💻 Written in Go, fully open-source

---

## 📦 Installation

Clone and build manually:

```bash
git clone https://github.com/LaviruD/grepr.git
cd grepr
go build -o grepr
````

Or download the prebuilt binary (coming soon 👀)

---

## 🧪 Usage

```bash
./grepr -i urls.txt -f js,php -r admin.*,login
```

### ⚙️ Available Flags

| Flag               | Description                                               |
| ------------------ | --------------------------------------------------------- |
| `-i, --input`      | **(Required)** Input file containing URLs                 |
| `-o, --output`     | Output file name (default: `output.txt`)                  |
| `-f, --filetypes`  | Comma-separated filetypes to filter (e.g., `js,php,aspx`) |
| `-r, --regex-list` | Regex patterns (comma-separated, e.g., `admin.*,login`)   |
| `--regex-file`     | File path containing regex patterns (one per line)        |
| `-s, --soora`      | Enable **Soora Super Mode** for deep filtering            |
| `-n, --nobanner`   | Disable the startup banner                                |

---

## 🧠 Soora Super Mode

Soora is a built-in, intelligent mode that:

* Filters common sensitive filetypes
* Applies regex patterns for juicy paths
* Searches for known keywords like `admin`, `login`, `config`, etc.

```bash
./grepr -i all-urls.txt -s
```

All filtered results will be saved as multiple intermediate files and a final merged output:
`Final-Grepr.txt`

---

## 📂 Output Structure

Each filter writes a separate file:

* `output-filetypes-Grepr.txt`
* `output-regexes-Grepr.txt`
* `Final-Grepr.txt` *(when Soora mode is used)*

Each file output includes:

* ✅ Total matched lines
* 📦 Output file size (KB)

---

## 🖥 Example

```bash
./grepr -i subdomains.txt -f js,php -r admin.*,login
```

Output:

```
[✓] Filetype filtered results written to: output-filetypes-Grepr.txt (37 lines, 12.89 KB)
[✓] Multi-regex filtered lines written to: output-regexes-Grepr.txt (11 lines, 4.21 KB)
```

---

## 🔧 Developer Info

* 👨‍💻 Developed by: Laviru Dilshan  
* 🌐 GitHub: [github.com/LaviruD](https://github.com/LaviruD)  
* 💼 LinkedIn: [linkedin.com/in/laviru-dilshan](https://www.linkedin.com/in/lavirudev)  
* 🐦 X (Twitter): [x.com/laviru_dilshan](https://x.com/laviru_dev)  

---

## 🛡️ License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/LaviruD/grepr/blob/main/LICENSE) file for details.

---

## 💬 Feedback

Found a bug or have an idea?
Open an issue or reach out via socials. Contributions are welcome!

---
