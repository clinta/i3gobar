package i3gobar

import (
	"fmt"
	"time"

	"github.com/distatus/battery"
)

func Batt(uc chan<- []I3Block) {
	b := make([]I3Block, 2)
	b[0].FullText = "Bat:"
	b[0].NoSeparator = true
	b[0].SeparatorBlockWidth = 3

	for {
		batteries, err := battery.GetAll()
		if err != nil {
			fmt.Println("Could not get battery info!")
			return
		}
		for i, bat := range batteries {
			b[i+1].FullText = fmt.Sprintf("%3.0f", (bat.Current/bat.Full)*100)
			b[i+1].Color = GetColor(1 - (bat.Current / bat.Full))
			if i != len(batteries)-1 {
				b[i+1].SeparatorBlockWidth = 3
				b[i+1].NoSeparator = true
			}
		}
		uc <- b

		time.Sleep(10 * time.Second)
	}
}
