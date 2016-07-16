package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type i3Block struct {
	FullText            string `json:"full_text"`
	ShortText           string `json:"short_text,omitempty"`
	Color               string `json:"color,omitempty"`
	Background          string `json:"background,omitempty"` // background color
	Border              string `json:"border,omitempty"`     // border color
	MinWidth            int    `json:"min_width,omitempty"`
	Align               string `json:"align,omitempty"` // center, right or left alignment of text in block
	Name                string `json:"name,omitempty"`  // not used by i3, used for click_events
	Instance            string `json:"instance,omitempty"`
	Urgent              bool   `json:"urgent,omitempty"`
	Separator           bool   `json:"separator,omitempty"`             // separator after the block
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"` // pixels to be left blank after the block
	Markup              string `json:"markup,omitempty"`                // set to pango for pango markup
}

type update struct {
	index  int
	update i3Block
}

var l *log.Logger

func printStatus() {
	l = log.New(os.Stderr, "", 0)

	pfuncs := [...]func(chan<- i3Block){
		cpus,
	}

	uc := make(chan update)
	for i, pfunc := range pfuncs {
		go func() {
			bc := make(chan i3Block)
			go pfunc(bc)
			for {
				b := <-bc
				uc <- update{
					index:  i,
					update: b,
				}
			}
		}()
	}

	fmt.Println("[")
	blocks := make([]i3Block, len(pfuncs))
	for {
		u := <-uc
		blocks[u.index] = u.update

		blockJson, err := jsonMarshal(blocks)
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Print(string(blockJson))
		fmt.Println(",")
	}
}

func getColor(n interface{}) string {
	// #00FF00
	i := 0
	s := fmt.Sprintf("%.0f", n)
	if pi, err := strconv.Atoi(s); err == nil {
		i = pi
	}

	r := (255 * i) / 100
	g := (255 * (100 - i)) / 100
	b := 0
	return fmt.Sprintf("#%0.2x%0.2x%0.2x", r, g, b)
}

func jsonMarshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)

	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)

	return b, err
}
