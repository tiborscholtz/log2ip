# log2ip

## Overview

**log2ip** is a cybersecurity-focused CLI tool written in Go that collects and analyzes IP address–related log entries from system log files.  
It is designed to help security engineers and system administrators quickly inspect authentication and SSH-related activity from multiple log sources in a simplified, readable format.

## Features

- Collects IP-related log entries from:
  - `auth.log`
  - `openssh` logs
  - Other supported system log files
- CLI-based **table workflow** for easy inspection
- Log message **simplification** for readability

## Log Simplification

Verbose log entries are normalized into concise descriptions.

**Example:**

Original log entry:  
```text
pam_unix(sshd:auth): authentication failure; logname= uid=0 euid=0 tty=ssh ruser= rhost=103.99.0.122
```

Simplified output:  
```text
authentication failure
````

## Usage

```bash
log2ip
````

## Installation

### Build from source

```bash
git clone https://github.com/tiborsholtz/log2ip.git
cd log2ip
Make build
```

## Requirements

* Go 1.20+
* Linux-based system with standard authentication logs

## Project Structure

```
log2ip/
├── bin/
├── cmd/log2ip/
├── internal/
├── internal/cache
├── internal/components
├── internal/config
├── internal/domain
├── internal/parsers
├── internal/resources
├── internal/ui
├── scripts/
├── testdata/
├── go.mod
├── go.sum
├── Makefile
├── CONTRIBUTING.md
└── README.md
```

## Security Use Cases

* Detect repeated authentication failures
* Identify suspicious IP addresses
* Review SSH access attempts

## Completed features

* [x] Pagination

## Roadmap

* [ ] Sort by columns
* [ ] Search in available columns
* [ ] Additional log support
* [ ] Pre-filter using command line parameters
* [ ] Auto-recognize log file type
* [ ] Export to CSV / JSON
* [ ] IP reputation lookups
* [ ] Time-range filtering
* [ ] Colorized output