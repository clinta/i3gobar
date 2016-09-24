package i3gobar

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	pango = "pango"
)

// I3Block represents a block to be printed on the bar. See the i3bar protocol documentation for details on each property
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
	Urgent              bool   `json:"urgent"`
	NoSeparator         bool   `json:"-"`                               // Specify as true to omit the spearator after this field
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"` // pixels to be left blank after the block
	Markup              string `json:"markup,omitempty"`                // set to pango for pango markup
}

type i3Block struct {
	I3Block
	Separator bool `json:"separator"` // separator after the block
}

type i3Header struct {
	Version int `json:"version"`
	// StopSignal  int  `json:"stop_signal"`
	// ContSignal  int  `json:"cont_signal"`
	ClickEvents bool `json:"click_events"`
}

type update struct {
	index  int
	update []I3Block
}

var logger *log.Logger

// Run runs all the specified functions, and prints the output to be consumed by i3bar.
// It runs each function in a goroutine and updates the bar when any of them return data on the return channel.
func Run(f []func(chan<- []I3Block), noSeparator bool, separatorBlockWidth int) {
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
		go func(i int, p func(chan<- []I3Block)) {
			bc := make(chan []I3Block)
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
	blocks := make([][]I3Block, len(f))
	for {
		u := <-uc
		blocks[u.index] = u.update

		var eblocks []I3Block
		for _, b := range blocks {
			if len(b) > 0 {
				b[len(b)-1].NoSeparator = noSeparator
				b[len(b)-1].SeparatorBlockWidth = separatorBlockWidth
			}
			eblocks = append(eblocks, b...)
		}

		blockJSON, err := marshallI3Block(eblocks)
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Print(string(blockJSON))
		fmt.Println(",")
	}
}

func marshallI3Block(bs []I3Block) ([]byte, error) {
	is := make([]i3Block, len(bs))
	for i, b := range bs {
		is[i] = i3Block{
			I3Block:   b,
			Separator: !b.NoSeparator,
		}
		//is[i].Separator = !is[i].NoSeparator
	}
	return jsonMarshal(is)
}

func jsonMarshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)

	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)

	return b, err
}

// GetColor returns a color between green and red where 0 = green and 100 = red
func GetColor(n float64) string {
	if n > 1 {
		n = 1
	}
	// #00FF00
	r := int(255 * (n * 2))
	g := 255
	b := 0

	if r >= 255 {
		r = 255
		g = int(255 * ((1 - n) * 2))
	}

	if g > 255 {
		g = 255
	}

	return fmt.Sprintf("#%0.2x%0.2x%0.2x", r, g, b)
}

// ColorString returns a pango formatted string colored between green and red with the provided value
func ColorString(s string, n float64) string {
	return fmt.Sprintf("<span foreground=\"%v\">%v</span>", GetColor(n), s)
}

func readLine(path string) string {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	return scanner.Text()
}
