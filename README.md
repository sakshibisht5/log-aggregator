Log Aggregator & Shipper

A lightweight log aggregation service built in Go, capable of tailing log files, parsing new entries, and shipping them to a destination.

Overview

This project implements a simplified version of log collectors like Filebeat or Logstash. It continuously monitors log files in real-time, parses log lines, and ships them to stdout. The project demonstrates:

Real-time file tailing.

Concurrent log processing with Go channels.

Pluggable parser and shipper architecture.

Graceful handling of file rotation and long-running services.
Minimal external dependencies (hpcloud/tail)

Features

File Tailing

Uses hpcloud/tail to watch logs in real-time

Handles log file rotations gracefully

Concurrent Pipeline

Processes log lines concurrently using Go channels

Separates parsing and shipping logic for modularity

Pluggable Shippers

Current implementation: stdout shipping

Can be extended to file, HTTP, or TCP endpoints

Logging Utilities

Minimal logging wrapper for info/error messages

Configuration

Supports YAML configuration (paths, sink type)

Easy to extend for multiple sources or sinks

Folder Structure
log-project/
├─ cmd/loggie/       # Main application
├─ pkg/core/log/     # Logging utilities
├─ pkg/source/file/  # File tailing
├─ pkg/sink/stdout/  # Stdout shipping
├─ conf/loggie.yaml  # Config file
├─ test.log          # Sample log file
├─ go.mod
├─ README.md

Getting Started
1. Clone and Setup
git clone https://github.com/<your-username>/log-aggregator.git
cd log-aggregator
go mod tidy

2. Run the Log Aggregator
cd cmd/loggie
go run main.go

3. Test Logging

Append lines to the test log:

echo "New log line" >> ../../test.log


Lines will appear in stdout in real-time.

Technologies Used

Go – Concurrency and system programming

hpcloud/tail – File tailing

Git & GitHub – Version control

YAML – Configuration management

Learning Outcomes

Advanced file I/O and filesystem event watching

Designing concurrent data pipelines with Go channels

Building robust long-running services

Parsing and structuring log data

Graceful shutdown and resource management

Future Enhancements:

Add multiple pluggable shippers. (HTTP, TCP, file)

Implement JSON and key-value parsers.

Add metrics and observability. (Prometheus)

Support multiple log sources with dynamic configuration.
