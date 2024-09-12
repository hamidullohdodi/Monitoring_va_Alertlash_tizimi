package monitor

//func CheckSystemStatus(config *config.Config) error {
//	// CPU ishlatilishini olish
//	cpuUsageStr, err := GetCPUUsage()
//	if err != nil {
//		return fmt.Errorf("CPU ma'lumotlarini olishda xato: %v", err)
//	}
//
//	// Memory ishlatilishini olish
//	memoryUsage, err := GetMemoryUsage()
//	if err != nil {
//		return fmt.Errorf("Memory ma'lumotlarini olishda xato: %v", err)
//	}
//
//	// Disk I/O ishlatilishini olish
//	diskIOUsage, diskIOUsage1, err := GetDiskIOUsage()
//	if err != nil {
//		return fmt.Errorf("Disk ma'lumotlarini olishda xato: %v", err)
//	}
//
//	// Agar threshold qiymatlaridan oshsa, logga yozish
//	if cpuUsageStr > config.CPU_THRESHOLD {
//		log.Printf("CPU ishlatilishi threshold'dan oshdi: %.2f%%", cpuUsageStr)
//	}
//	if memoryUsage > config.MEMORY_THRESHOLD {
//		log.Printf("Xotira ishlatilishi threshold'dan oshdi: %d%%", memoryUsage)
//	}
//	if diskIOUsage > config.DISK_IO_THRESHOLD {
//		log.Printf("Disk I/O ishlatilishi threshold'dan oshdi:  oqish->%f kB and yozish->%f kB", diskIOUsage, diskIOUsage1)
//	}
//
//	return nil
//}
