package i3gobar

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/mem"
)

// MemFree prints the free memory in human readable units. Is green as long as there is 1GiB free, shifts red as the free memory approaches zero.
func MemFree(uc chan<- I3Block) {
	var o I3Block
	o.Markup = pango

	for {
		vm, err := mem.VirtualMemory()
		if err != nil {
			logger.Println(err)
			continue
		}

		memc := float64(0)
		if vm.Available < 1024^3 {
			memc = float64((1<<24 - vm.Available) / 1 << 24)
		}
		memStats := ColorString(humanize.IBytes(vm.Available), memc)
		o.FullText = fmt.Sprintf("Free: %v", memStats)

		uc <- o

		time.Sleep(1 * time.Second)
	}
}

// SwapUsed shows the swap space in use. Turns red once more than 100 bytes are used.
func SwapUsed(uc chan<- I3Block) {
	var o I3Block
	o.Markup = pango

	for {
		sm, err := mem.SwapMemory()
		if err != nil {
			logger.Println(err)
			continue
		}
		swStats := ColorString(humanize.IBytes(sm.Used), float64(sm.Used))
		o.FullText = fmt.Sprintf("Swap: %v", swStats)

		uc <- o

		time.Sleep(1 * time.Second)
	}
}
