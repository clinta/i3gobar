package i3gobar

import (
	"time"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/mem"
)

// MemFree prints the free memory in human readable units. Is green as long as there is 1GiB free, shifts red as the free memory approaches zero.
func MemFree(uc chan<- []I3Block) {
	b := make([]I3Block, 2)
	b[0].FullText = "Free:"
	b[0].NoSeparator = true

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
		b[1].Color = GetColor(memc)
		b[1].FullText = humanize.IBytes(vm.Available)

		uc <- b

		time.Sleep(1 * time.Second)
	}
}

// SwapUsed shows the swap space in use. Turns red once more than 100 bytes are used.
func SwapUsed(uc chan<- []I3Block) {
	b := make([]I3Block, 2)
	b[0].FullText = "Swap:"
	b[0].NoSeparator = true

	for {
		sm, err := mem.SwapMemory()
		if err != nil {
			logger.Println(err)
			continue
		}
		swc := float64(sm.Used / (1 << 24))
		b[1].Color = GetColor(swc)
		b[1].FullText = humanize.IBytes(sm.Used)

		uc <- b

		time.Sleep(1 * time.Second)
	}
}
