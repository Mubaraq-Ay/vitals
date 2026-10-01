package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
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

}
