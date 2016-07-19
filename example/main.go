package main

import (
	"time"

	"github.com/clinta/i3gobar"
)

func main() {
	f := []func(chan<- []i3gobar.I3Block){
		i3gobar.LoadAvg,
		i3gobar.CPU,
		i3gobar.MemFree,
		i3gobar.SwapUsed,
		i3gobar.DateTime,
		easyFunction,
	}

	i3gobar.Run(f)
}

func easyFunction(uc chan<- []i3gobar.I3Block) {
	b := make([]i3gobar.I3Block, 1)
	o := &b[0]
	for {
		o.FullText = i3gobar.ColorString("Super easy!", 0)
		o.Markup = "pango"
		uc <- b
		time.Sleep(60 * time.Second)
	}
}
