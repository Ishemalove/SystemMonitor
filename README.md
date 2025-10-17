# 🧠 Windows System Optimizer (WinOptimizer)

## 📘 Project Description

WinOptimizer is a powerful Windows-compatible system monitoring dashboard built purely in Go (Golang). It provides real-time insights into CPU usage, memory consumption, and disk utilization using the `gopsutil` library. This project helps developers learn system-level performance monitoring and backend dashboard development using Go’s concurrency and performance capabilities.

Short description (for project input box):

> A Go-powered Windows system dashboard that monitors CPU, memory, and disk usage in real-time.

---

## 🚀 Features

* Real-time CPU usage monitoring (auto-updates)
* Memory (RAM) usage visualization
* Disk usage and available space tracking
* Local web dashboard served via Go’s built-in HTTP server
* Designed specifically for Windows
* Extendable for logs and system alerts

---

## 🛠️ Installation & Setup

### Prerequisites

* Go 1.21+ installed
* Windows system

### Steps

```bash
# 1. Clone the repository
git clone https://github.com/yourusername/winoptimizer.git
cd winoptimizer

# 2. Initialize Go module
go mod tidy

# 3. Run the project
go run main.go
```

### Open in Browser

```
http://localhost:8080
```

View your live dashboard directly in your browser.

---

## 🧩 Project Structure

```
winoptimizer/
│
├── main.go           # Main application entry point
├── go.mod            # Go module definition
├── go.sum            # Dependency checksums
└── README.md         # Project documentation
```

---

## 🖼️ Screenshots

### 1️⃣ Dashboard Overview

![Dashboard Overview](https://via.placeholder.com/800x400?text=System+Dashboard+Overview)
Displays CPU, memory, and disk statistics updated in real time.

### 2️⃣ CPU Usage Graph

![CPU Graph](https://via.placeholder.com/800x400?text=CPU+Usage+Graph)
Visual representation of system performance.

### 3️⃣ Memory and Disk View

![Memory and Disk](https://via.placeholder.com/800x400?text=Memory+and+Disk+Usage)
Tracks total, used, and free memory and storage.

---

## 📚 Technologies Used

* **Golang** — main programming language
* **gopsutil** — access system metrics
* **net/http** — serves the dashboard locally

---

## 🧠 Learning Objectives

* Use Go for real-time system monitoring
* Manage Go dependencies and modules
* Serve live data over HTTP
* Learn concurrency and system API integration

---

## 🔮 Future Improvements

* Add performance logs and alerts
* Export data as CSV or JSON
* Integrate network monitoring
* Add modern UI using Go templates or React

---

## 👩‍💻 Author

**ISHEMA NKERABAHIZI Love**
Rwanda Coding Academy
A passionate developer creating efficient and educational software in Go.

---

## 🏁 License

This project is open-source under the MIT License.
