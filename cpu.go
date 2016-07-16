package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func cpus(uc chan<- i3Block) {
	var o i3Block
	o.Markup = "pango"
	for {
		percs, err := cpu.Percent(0, true)
		if err != nil {
			l.Println(err)
			continue
		}

		o.FullText = "CPU:"

		for _, p := range percs {
			s := fmt.Sprintf("<span foreground=\"%v\">%3.0f</span>", getColor(p), p)
			//s := fmt.Sprintf("%3.0f", p)
			o.FullText = fmt.Sprintf("%v%v", o.FullText, s)
		}
		uc <- o

		time.Sleep(1 * time.Second)
	}
}
