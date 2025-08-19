# **Log Aggregator & Shipper**

**A lightweight log aggregation service built in Go, capable of tailing log files, parsing new entries, and shipping them to a destination.**

---

## **Overview**

This project implements a simplified version of log collectors like **Filebeat** or **Logstash**.  
It continuously monitors log files in real-time, parses log lines, and ships them to **stdout**.

**Demonstrates:**

- Real-time file tailing
- Concurrent log processing with Go channels
- Pluggable parser and shipper architecture
- Graceful handling of file rotation
- Minimal external dependencies (`hpcloud/tail`)

---

## **Features**

### **1. File Tailing**
- Real-time monitoring using `hpcloud/tail`
- Handles log file rotations gracefully

### **2. Concurrent Pipeline**
- Uses **Go channels** for concurrency
- Separates parsing and shipping logic

### **3. Pluggable Shippers**
- Current implementation: `stdout`
- Easily extendable to file, HTTP, or TCP

### **4. Logging Utilities**
- Minimal logging wrapper for **info/error messages**

### **5. Configuration**
- Supports **YAML configuration** (`paths`, `sink`)
- Easy to extend for multiple sources or sinks

---

## **Folder Structure**

log-project/
 cmd/loggie/ # Main application

pkg/core/log/ # Logging utilities

pkg/source/file/ # File tailing

pkg/sink/stdout/ # Stdout shipping

conf/loggie.yaml # Config file

test.log # Sample log file

go.mod

README.md

- Run the Log Aggregator
- Test Logging

## **Technologies Used**

- **Go** – Concurrency and system programming
- **hpcloud/tail** – File tailing
- **Git & GitHub** – Version control
- **YAML** – Configuration management

---

## **Learning Outcomes**

- Advanced file I/O and filesystem event watching
- Designing concurrent data pipelines with Go channels
- Building robust long-running services
- Parsing and structuring log data
- Graceful shutdown and resource management

---

## **Future Enhancements**

- Add multiple pluggable shippers (HTTP, TCP, file)
- Implement JSON and key-value parsers
- Add metrics and observability (Prometheus)
- Support multiple log sources with dynamic configuration

