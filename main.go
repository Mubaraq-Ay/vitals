package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)
 
func main() {
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
	fmt.Printf("CPU: %.2f%%\n", percent[0])
	fmt.Printf("Total memory: %d\n", memory.Total)

	usage, err := disk.Usage("C:\\")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Disk path: %s\n", usage.Path)
	fmt.Printf("Total: %d bytes\n", usage.Total)
	fmt.Printf("Free: %d bytes\n", usage.Free)
	fmt.Printf("Used: %.2f%%\n", usage.UsedPercent)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hostname:", hostname)

	

}