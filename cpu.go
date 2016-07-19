package i3gobar

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// CPU returns cpu time usage every second for each core. Color shifts from green to red as the usage approaches 100
func CPU(uc chan<- []I3Block) {
	b := make([]I3Block, 1)
	b[0].FullText = "CPU:"
	b[0].NoSeparator = true
	b[0].SeparatorBlockWidth = 3

	for {
		percs, err := cpu.Percent(0, true)
		if err != nil {
			logger.Println(err)
			continue
		}
		if len(b) < len(percs)+1 {
			b = append(b[:1], make([]I3Block, len(percs))...)
		}

		for i, p := range percs {
			b[i+1].FullText = fmt.Sprintf("%3.0f", p)
			b[i+1].Color = GetColor(p / 100)
			if i != len(percs)-1 {
				b[i+1].SeparatorBlockWidth = 3
				b[i+1].NoSeparator = true
			}
		}
		uc <- b

		time.Sleep(1 * time.Second)
	}
}
