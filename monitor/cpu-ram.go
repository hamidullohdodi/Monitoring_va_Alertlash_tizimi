package monitor

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var (
	lastTotalTime  int
	lastActiveTime int
	lastReadTime   time.Time
)

func GetCPUUsage() (float64, error) {
	data, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return 0, fmt.Errorf("CPU ma'lumotlarini o'qishda xato: %v", err)
	}

	fields := strings.Fields(string(data))
	if len(fields) < 0 {
		return 0, fmt.Errorf("Noto'g'ri format")
	}

	userTime, _ := strconv.Atoi(fields[1])
	systemTime, _ := strconv.Atoi(fields[3])
	idleTime, _ := strconv.Atoi(fields[4])

	totalTime := userTime + systemTime + idleTime
	activeTime := userTime + systemTime

	if !lastReadTime.IsZero() {
		timeDelta := time.Since(lastReadTime).Seconds()
		totalDelta := float64(totalTime - lastTotalTime)
		activeDelta := float64(activeTime - lastActiveTime)

		if totalDelta > 0 && timeDelta > 0 {
			cpuUsage := (activeDelta / totalDelta) * 100
			lastTotalTime = totalTime
			lastActiveTime = activeTime
			lastReadTime = time.Now()
			return cpuUsage, nil
		}
	}

	lastTotalTime = totalTime
	lastActiveTime = activeTime
	lastReadTime = time.Now()

	return 0.00, nil
}
