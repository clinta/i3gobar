package i3gobar

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func CPUGraph(uc chan<- []I3Block) {
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

		char := make([]string, 8)
		char[0] = "\u2581"
		char[1] = "\u2582"
		char[2] = "\u2583"
		char[3] = "\u2584"
		char[4] = "\u2585"
		char[5] = "\u2586"
		char[6] = "\u2587"
		char[7] = "\u2588"
		for i, p := range percs {
			b[i+1].FullText = fmt.Sprintf("%v", char[int((p/100)*7)])
			//b[i+1].FullText = fmt.Sprintf("%3.0f", p)
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

func CPUTemp(uc chan<- []I3Block) {
	b := make([]I3Block, 1)
	for {
		temp, _ := strconv.Atoi(readLine("/sys/class/thermal/thermal_zone0/temp"))
		temp = temp / 1000
		b[0].FullText = fmt.Sprintf("%v \u2103", temp)
		//color range 30 to 65 celcius
		base := temp - 30
		if base < 0 {
			base = 0
		}
		b[0].Color = GetColor(float64(base) / 35)
		time.Sleep(1 * time.Second)
		uc <- b
	}
}
