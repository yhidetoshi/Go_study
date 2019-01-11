package main

import (
	"fmt"
	"github.com/cloudfoundry/gosigar"
)

// MB
func format(val uint64) uint64 {
	return val / 1024 / 1024
}

func main(){
	uptime := sigar.Uptime{}
	uptime.Get()

	// UPTIME
	fmt.Printf("Uptime: %s \n", uptime.Format())

	avg := sigar.LoadAverage{}
	avg.Get()


	// Load Average
	fmt.Printf("Load Average %.2f, %.2f \n", avg.One, avg.Five)

	mem := sigar.Mem{}
	mem.Get()

	// Memory Usage[MB]
	fmt.Printf("Memory[MB] total=%d, used=%d, free=%d \n", format(mem.Total), format(mem.Used), format(mem.Free))

}

/* 実行結果
Uptime: 15 days, 14:05
Load Average 2.09, 2.49
Memory[MB] total=8192, used=8112, free=79
*/