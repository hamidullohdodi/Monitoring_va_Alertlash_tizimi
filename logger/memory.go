package logger

import (
	"dodi/bot"
	"github.com/shirou/gopsutil/process"
	"log"
	"os"
	"sort"
)

type ProcInfo1 struct {
	Name  string
	Usage float64
}

type ByUsage1 []ProcInfo1

func (a ByUsage1) Len() int      { return len(a) }
func (a ByUsage1) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByUsage1) Less(i, j int) bool {
	return a[i].Usage > a[j].Usage
}

func Memory() {
	file, err := os.OpenFile("memory_usage.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Fayl ochishda xato: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	for {
		processes, _ := process.Processes()

		var procinfos []ProcInfo1
		for _, p := range processes {
			a, _ := p.CPUPercent()
			n, _ := p.Name()
			procinfos = append(procinfos, ProcInfo1{n, a})
		}
		sort.Sort(ByUsage1(procinfos))

		log.Println("Top 5 processes by memory usage:")
		for _, p := range procinfos[:5] {
			if p.Usage > 0 {
				bot.BotLog("Memory ishlatilishi threshold'dan oshdi:  shu ilovada -> %s " + p.Name + " shunchaga oshib ketdi " + Float64ToString(p.Usage))
			}
			log.Printf("   %s -> %f", p.Name, p.Usage)
		}
		return
	}
}
