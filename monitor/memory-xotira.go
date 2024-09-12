package monitor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	lastTotalMem    int
	lastFreeMem     int
	lastMemReadTime time.Time
)

func GetMemoryUsage() (float64, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("Xotira ma'lumotlarini o'qishda xato: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalMem := 0
	freeMem := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
			totalMem, _ = parseMemoryLine(line)
		}
		if strings.HasPrefix(line, "MemFree:") {
			freeMem, _ = parseMemoryLine(line)
		}
	}

	if totalMem == 0 {
		return 0, fmt.Errorf("Xotira ma'lumotlari mavjud emas")
	}

	usedMem := totalMem - freeMem

	memUsage := float64(usedMem) / float64(totalMem) * 100

	if !lastMemReadTime.IsZero() {
		timeDelta := time.Since(lastMemReadTime).Seconds()
		totalDelta := float64(totalMem - lastTotalMem)
		freeDelta := float64(freeMem - lastFreeMem)

		if totalDelta > 0 && timeDelta > 0 {
			memUsage := (totalDelta - freeDelta) / totalDelta * 100
			lastTotalMem = totalMem
			lastFreeMem = freeMem
			lastMemReadTime = time.Now()
			return memUsage, nil
		}
	}

	lastTotalMem = totalMem
	lastFreeMem = freeMem
	lastMemReadTime = time.Now()

	return memUsage, nil
}

func parseMemoryLine(line string) (int, error) {
	parts := strings.Fields(line)
	if len(parts) > 0 {
		return strconv.Atoi(parts[1])
	}
	return 0, fmt.Errorf("Xotira satri noto'g'ri formatda")
}
