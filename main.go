package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/disk"
    "time"
)

// Template for dashboard HTML
const tpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Windows System Performance Dashboard</title>
    <meta http-equiv="refresh" content="5"> <!-- auto refresh every 5s -->
    <style>
        body { font-family: Arial, sans-serif; background: #0d1117; color: #f0f6fc; padding: 20px; }
        h1 { color: #58a6ff; }
        .card { background: #161b22; padding: 15px; margin: 15px 0; border-radius: 10px; }
    </style>
</head>
<body>
    <h1>System Performance Dashboard</h1>
    <div class="card">
        <h3>CPU Usage</h3>
        <p>{{.CPU}}%</p>
    </div>
    <div class="card">
        <h3>Memory Usage</h3>
        <p>{{.Memory}}%</p>
    </div>
    <div class="card">
        <h3>Disk Usage</h3>
        <p>{{.Disk}}%</p>
    </div>
    <p><small>Last Updated: {{.Time}}</small></p>
</body>
</html>
`

// Struct to pass data to template
type Stats struct {
    CPU    float64
    Memory float64
    Disk   float64
    Time   string
}

// Handler for root route
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    // Get CPU usage
    cpuPercent, _ := cpu.Percent(time.Second, false)
    memStat, _ := mem.VirtualMemory()
    diskStat, _ := disk.Usage("C:/")

    stats := Stats{
        CPU:    cpuPercent[0],
        Memory: memStat.UsedPercent,
        Disk:   diskStat.UsedPercent,
        Time:   time.Now().Format("15:04:05"),
    }

    tmpl, err := template.New("dashboard").Parse(tpl)
    if err != nil {
        log.Println("Template error:", err)
        return
    }

    tmpl.Execute(w, stats)
}

func main() {
    fmt.Println("🚀 Starting Windows System Dashboard at http://localhost:8080 ...")
    http.HandleFunc("/", dashboardHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
