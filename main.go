package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)
 

func main() {
	percent, err := cpu.Percent(time.Second, false)

	if err != nil {
			fmt.Println(err)
			return
	}
	fmt.Println(percent[0]) 
	fmt.Printf("CPU: %.2f%%\n", percent[0])
}
