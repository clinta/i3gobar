package i3gobar

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

// LoadAvg prints the load average for 1, 5 and 15 minutes every second. Color for each load shifts from green to red as it approaches the number of cores on the system.
func LoadAvg(uc chan<- []I3Block) {
	b := make([]I3Block, 4)
	b[0].FullText = "Load:"
	b[0].NoSeparator = true
	//b[0].SeparatorBlockWidth = 3

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

		b[1].FullText = fmt.Sprintf("%01.02v", la.Load1)
		b[1].Color = GetColor(la.Load1 / cores)
		b[1].NoSeparator = true

		b[2].FullText = fmt.Sprintf("%01.02v", la.Load5)
		b[2].Color = GetColor(la.Load5 / cores)
		b[2].NoSeparator = true

		b[3].FullText = fmt.Sprintf("%01.02v", la.Load15)
		b[3].Color = GetColor(la.Load15 / cores)

		uc <- b

		time.Sleep(1 * time.Second)
	}
}
