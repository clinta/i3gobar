package i3gobar

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// CPU returns cpu time usage every second for each core. Color shifts from green to red as the usage approaches 100
func CPU(uc chan<- I3Block) {
	var o I3Block
	o.Markup = "pango"
	for {
		percs, err := cpu.Percent(0, true)
		if err != nil {
			logger.Println(err)
			continue
		}

		o.FullText = "CPU:"

		for _, p := range percs {
			s := fmt.Sprintf("<span foreground=\"%v\">%3.0f</span>", GetColor(p/100), p)
			//s := fmt.Sprintf("%3.0f", p)
			o.FullText = fmt.Sprintf("%v%v", o.FullText, s)
		}
		uc <- o

		time.Sleep(1 * time.Second)
	}
}
