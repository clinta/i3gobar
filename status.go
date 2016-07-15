package main

import (
	"encoding/json"
	"fmt"
)

type i3Block struct {
	FullText            string `json:"full_text"`
	ShortText           string `json:"short_text"`
	Color               string `json:"color"`
	Background          string `json:"background"` // background color
	Border              string `json:"border"`     // border color
	MinWidth            int    `json:"min_width"`
	Align               string `json:"align"` // center, right or left alignment of text in block
	Name                string `json:"name"`  // not used by i3, used for click_events
	Instance            string `json:"instance"`
	Urgent              bool   `json:"urgent"`
	Separator           bool   `json:"separator"`             // separator after the block
	SeparatorBlockWidth int    `json:"separator_block_width"` // pixels to be left blank after the block
	Markup              string `json:"markup"`                // set to pango for pango markup
}

type update struct {
	index  int
	update []i3Block
}

func printStatus() {
	pfuncs := []func(chan<- update){
		cpus,
	}

	uc = make(chan update)

	for i, pfunc := range pfpuncs {
		go pfunc(uc)
	}

	for {
		u := <-uc
	}
}
