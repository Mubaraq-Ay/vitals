# Vitals

A lightweight system diagnostics CLI built with Go.

Vitals gives you a quick overview of your machine's current system information and resource usage directly from the terminal.

## Features

* CPU usage
* Memory usage
* Disk usage
* Hostname
* Operating system
* System uptime
* Human-readable resource values
* Simple terminal output

## Example

```text
Vitals
======

System
------
Hostname: DESKTOP-XXXXX
OS: windows
Uptime: 5h 42m 18s

Disk
------
Path: C:\
Total: 476.01 GB
Free: 156.72 GB
Used: 67.09%

Memory
------
Total: 15.86 GB
Used: 8.21 GB
Available: 7.65 GB

CPU
------
Usage: 4.21%
```

## Requirements

* Go 1.27.1 or later

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

## Built With

* Go
* gopsutil

## Why Vitals?

Vitals started as a rewrite of an earlier system health checker I built in Python.

I built it as a small, practical way to learn Go by working with its standard library, external packages, error handling, types, formatting, and system-level information.

## License

This project is for learning and personal use.
