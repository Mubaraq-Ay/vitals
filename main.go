package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	fmt.Println("Vitals")
	fmt.Println("========")
	fmt.Println()

	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
		return
	}

	totalMemoryGB := float64(memory.Total) / 1024 / 1024 / 1024
	usedMemoryGB := float64(memory.Used) / 1024 / 1024 / 1024
	availableMemoryGB := float64(memory.Available) / 1024 / 1024 / 1024

	usage, err := disk.Usage("C:\\")
	if err != nil {
		fmt.Println(err)
		return
	}

	totalDiskGB := float64(usage.Total) / 1024 / 1024 / 1024
	freeDiskGB := float64(usage.Free) / 1024 / 1024 / 1024

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}

	platform := runtime.GOOS

	uptime, err := host.Uptime()
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Duration(uptime) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	// cli

	fmt.Println("System")
	fmt.Println("------")
	fmt.Println()

	fmt.Println("Hostname:", hostname)
	fmt.Println("OS:", platform)
	fmt.Printf("Uptime: %dh %dm %ds\n", hours, minutes, seconds)

	fmt.Println()
	fmt.Println("Disk")
	fmt.Println("------")

	fmt.Printf("Disk path: %s\n", usage.Path)
	fmt.Printf("Total: %.2f GB\n", totalDiskGB)
	fmt.Printf("Free: %.2f GB\n", freeDiskGB)
	fmt.Printf("Used: %.2f%%\n", usage.UsedPercent)

	fmt.Println()
	fmt.Println("Memory")
	fmt.Println("------")
	fmt.Printf("Total: %.2f GB\n", totalMemoryGB)
	fmt.Printf("Used: %.2f GB\n", usedMemoryGB)
	fmt.Printf("Available: %.2f GB\n", availableMemoryGB)

	fmt.Println()
	fmt.Println("CPU")
	fmt.Println("------")
	fmt.Printf("CPU: %.2f%%\n", percent[0])

}
