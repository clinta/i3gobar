package i3gobar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	pango = "pango"
)

type I3Block struct {
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

type i3Header struct {
	Version int `json:"version"`
	// StopSignal  int  `json:"stop_signal"`
	// ContSignal  int  `json:"cont_signal"`
	ClickEvents bool `json:"click_events"`
}

type update struct {
	index  int
	update I3Block
}

var logger *log.Logger

func Run(f []func(chan<- I3Block)) {
	logger = log.New(os.Stderr, "", 0)

	hdr := &i3Header{
		Version:     1,
		ClickEvents: false,
	}

	hdrb, err := json.Marshal(hdr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hdrb))

	uc := make(chan update)
	for i, p := range f {
		go func(i int, p func(chan<- I3Block)) {
			bc := make(chan I3Block)
			go p(bc)
			for {
				b := <-bc
				uc <- update{
					index:  i,
					update: b,
				}
			}
		}(i, p)
	}

	fmt.Println("[")
	blocks := make([]I3Block, len(f))
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

func jsonMarshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)

	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)

	return b, err
}

func GetColor(n float64) string {
	// #00FF00
	r := int(255 * n)
	g := int(255 * (1 - n))
	b := 0
	return fmt.Sprintf("#%0.2x%0.2x%0.2x", r, g, b)
}

func ColorString(s string, n float64) string {
	return fmt.Sprintf("<span foreground=\"%v\">%v</span>", GetColor(n), s)
}
