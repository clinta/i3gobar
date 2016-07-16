package main

import (
	"time"

	"github.com/clinta/i3gobar"
)

func main() {
	f := []func(chan<- i3gobar.I3Block){
		i3gobar.LoadAvg,
		i3gobar.CPU,
		i3gobar.MemFree,
		i3gobar.SwapUsed,
		easyFunction,
	}

	i3gobar.Run(f)
}

func easyFunction(uc chan<- i3gobar.I3Block) {
	var o i3gobar.I3Block
	for {
		o.FullText = "Super easy!"
		uc <- o
		time.Sleep(60 * time.Second)
	}
}
