package main

import (
	"fmt"
	"strings"

	cpu "github.com/triopium/core_stats/internal/cpu"
)

func main() {
	cpuLoads, err := cpu.GetCPULoad()
	if err != nil {
		fmt.Println("Error fetching CPU load:", err)
		return
	}

	// Build the output
	output := strings.Builder{}
	for _, load := range cpuLoads {
		output.WriteString(cpu.MapToDot(load))
	}
	fmt.Println(output.String())
}
