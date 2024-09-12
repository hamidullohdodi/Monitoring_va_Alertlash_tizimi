package logger

import (
	"dodi/bot"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"log"
	"os"
	"sort"
)

type ProcInfo struct {
	Name  string
	Usage float64
}

type ByUsage []ProcInfo

func (a ByUsage) Len() int      { return len(a) }
func (a ByUsage) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByUsage) Less(i, j int) bool {
	return a[i].Usage > a[j].Usage
}

func Cpu() {
	file, err := os.OpenFile("cpu_usage.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Fayl ochishda xato: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	for {
		processes, _ := process.Processes()

		var procinfos []ProcInfo
		for _, p := range processes {
			a, _ := p.CPUPercent()
			n, _ := p.Name()
			procinfos = append(procinfos, ProcInfo{n, a})
		}
		sort.Sort(ByUsage(procinfos))

		log.Println("Top 5 processes by CPU usage:")
		for _, p := range procinfos[:5] {
			if p.Usage > 0 {
				bot.BotLog("CPU ishlatilishi threshold'dan oshdi:  shu ilovada -> %s " + p.Name + " shunchaga oshib ketdi " + Float64ToString(p.Usage))
			}
			log.Printf("   %s -> %f", p.Name, p.Usage)
		}
		return
	}
}

func Float64ToString(value float64) string {
	return fmt.Sprintf("%.2f", value)
}
