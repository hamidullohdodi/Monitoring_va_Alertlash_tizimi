package main

import (
	"dodi/bot"
	"dodi/config"
	"dodi/logger"
	"dodi/monitor"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	cfg := config.NewConfig()

	logFile, err := os.OpenFile(cfg.LOG_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	go monitorCPU(cfg)
	go monitorMemory(cfg)
	go monitorDiskIO(cfg)

	select {} // Infinite select to keep the program running
}

func monitorCPU(cfg *config.Config) {
	lastAlert := time.Now()
	for {
		cpuUsage, err := monitor.GetCPUUsage()
		if err != nil {
			log.Printf("error getting CPU usage: %v", err)
			continue
		}

		if cpuUsage >= cfg.CPU_THRESHOLD {
			now := time.Now()
			if now.Sub(lastAlert) > 1*time.Minute {
				log.Printf("CPU ishlatilishi threshold'dan oshdi: %.2f%%", cpuUsage)
				logger.Cpu() // Log to file
				bot.BotLog("CPU ishlatilishi threshold'dan oshdi: " + Float64ToString(cpuUsage))
				lastAlert = now
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func monitorMemory(cfg *config.Config) {
	lastAlert := time.Now()
	for {
		memoryUsage, err := monitor.GetMemoryUsage()
		if err != nil {
			log.Printf("error getting memory usage: %v", err)
			continue
		}

		if memoryUsage >= cfg.MEMORY_THRESHOLD {
			now := time.Now()
			if now.Sub(lastAlert) > 1*time.Minute {
				log.Printf("Xotira ishlatilishi threshold'dan oshdi: %.2f%%", memoryUsage)
				logger.Memory() // Log to file
				bot.BotLog("Xotira ishlatilishi threshold'dan oshdi: " + Float64ToString(memoryUsage))
				lastAlert = now
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func monitorDiskIO(cfg *config.Config) {
	lastAlert := time.Now()
	for {
		diskIOUsage, _, err := monitor.GetDiskIOUsage()
		if err != nil {
			log.Printf("error getting disk I/O usage: %v", err)
			continue
		}

		if diskIOUsage >= cfg.DISK_IO_THRESHOLD {
			now := time.Now()
			if now.Sub(lastAlert) > 1*time.Minute {
				log.Printf("Disk I/O ishlatilishi threshold'dan oshdi: %.2f%%", diskIOUsage)
				bot.BotLog("Disk I/O ishlatilishi threshold'dan oshdi: " + Float64ToString(diskIOUsage))
				lastAlert = now
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func Float64ToString(value float64) string {
	return fmt.Sprintf("%.2f", value)
}
