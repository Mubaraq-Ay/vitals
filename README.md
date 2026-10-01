# Vitals

A lightweight system diagnostics CLI written in Go.

Vitals provides a simple way to inspect basic system information and resource usage directly from the terminal.

## Features

- CPU usage
- Memory information
- Disk usage
- Hostname
- Operating system information
- System uptime

## Example

```text
Vitals

CPU: 2.15%
Total memory: 17019527168 bytes

Disk path: C:\
Total: 511110590464 bytes
Free: 168234340352 bytes
Used: 67.09%

Hostname: DESKTOP-XXXXX
OS: windows
```

## Requirements

- Go 1.XX or later

## Installation

Clone the repository:

```bash
git clone https://github.com/Mubaraq-Ay/vitals.git
cd vitals
```

Install dependencies:

```bash
go mod tidy
```

Run Vitals:

```bash
go run .
```

## Dependencies

Vitals uses [gopsutil](https://github.com/shirou/gopsutil) to retrieve system and hardware information.

## Why Vitals?

Vitals is a Go rewrite of an earlier system health checker I built in Python.

The goal is to explore Go by building something small and practical while learning the language, standard library, error handling, packages, and system-level information.

## Status

Work in progress.

More system diagnostics and cleaner output will be added as development continues.

## License

This project is for learning and personal use.