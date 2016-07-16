package i3gobar

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

// LoadAvg prints the load average for 1, 5 and 15 minutes every second. Color for each load shifts from green to red as it approaches the number of cores on the system.
func LoadAvg(uc chan<- I3Block) {
	var o I3Block
	o.Markup = pango

	c, err := cpu.Counts(false)
	if err != nil {
		logger.Println(err)
		c = 1
	}
	cores := float64(c)

	for {
		la, err := load.Avg()
		if err != nil {
			logger.Println(err)
			continue
		}

		l1 := ColorString(fmt.Sprintf("%01.02v", la.Load1), la.Load1/cores)
		l5 := ColorString(fmt.Sprintf("%01.02v", la.Load5), la.Load5/cores)
		l15 := ColorString(fmt.Sprintf("%01.02v", la.Load15), la.Load15/cores)

		o.FullText = fmt.Sprintf("Load: %v %v %v", l1, l5, l15)

		uc <- o

		time.Sleep(1 * time.Second)
	}
}
