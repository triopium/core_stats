package cpu

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Characters to represent load levels (stacked dots)
var dots = []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}

// Colors for different load levels
var colors = []string{"#00FF00", "#55FF00", "#AAFF00", "#FFFF00", "#FFAA00", "#FF5500", "#FF0000", "#FF0000"}

// Fetch per-core CPU load using mpstat
func GetCPULoad() ([]float64, error) {
	cmd := exec.Command("mpstat", "-P", "ALL", "1", "1")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	cpuLoads := []float64{}

	// Regular expression to match CPU lines
	re := regexp.MustCompile(`^\s*Average:\s+(\d+)\s+.*\s+([\d.]+)$`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 3 {
			idle, err := strconv.ParseFloat(match[2], 64)
			if err == nil {
				cpuLoads = append(cpuLoads, 100-idle)
			}
		}
	}

	return cpuLoads, nil
}

// Map CPU usage to a dot and color
func MapToDot(usage float64) string {
	level := int(usage / 12.5)
	if level >= len(dots) {
		level = len(dots) - 1
	}
	return fmt.Sprintf("<span color=\"%s\">%s</span>", colors[level], dots[level])
}
