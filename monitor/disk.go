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
	lastDiskReadBytes  int
	lastDiskWriteBytes int
	lastDiskReadTime   time.Time
)

func GetDiskIOUsage() (float64, float64, error) {
	file, err := os.Open("/proc/diskstats")
	if err != nil {
		fmt.Println("lllllllllllllllllllllllllllllllll11111111")
		return 0, 0, fmt.Errorf("Disk I/O ma'lumotlarini o'qishda xato: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readBytes := 0
	writeBytes := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 14 {
			read, _ := strconv.Atoi(parts[5])
			write, _ := strconv.Atoi(parts[9])
			readBytes += read
			writeBytes += write
		}
	}

	if lastDiskReadTime.IsZero() {
		lastDiskReadBytes = readBytes
		lastDiskWriteBytes = writeBytes
		lastDiskReadTime = time.Now()
		fmt.Println("lllllllllllllllllllllllllllllllll222222222222222222")
		return float64(readBytes / 1000), float64(writeBytes / 1000), nil
	}

	timeDelta := time.Since(lastDiskReadTime).Seconds()

	readDelta := float64(readBytes - lastDiskReadBytes)
	writeDelta := float64(writeBytes - lastDiskWriteBytes)

	if timeDelta > 0 {
		readRate := readDelta / timeDelta / 1000
		writeRate := writeDelta / timeDelta / 1000

		lastDiskReadBytes = readBytes
		lastDiskWriteBytes = writeBytes
		lastDiskReadTime = time.Now()

		fmt.Println("lllllllllllllllllllllllllllllllll33333333333333")
		return readRate, writeRate, nil
	}
	fmt.Println("lllllllllllllllllllllllllllllllll4444444444444")

	return float64(readBytes / 1000), float64(writeBytes / 1000), nil
}
