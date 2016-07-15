package main

import (
	"encoding/json"
	"fmt"
)

type i3Header struct {
	Version int `json:"version"`
	// StopSignal  int  `json:"stop_signal"`
	// ContSignal  int  `json:"cont_signal"`
	ClickEvents bool `json:"click_events"`
}

func main() {
	hdr := &i3Header{
		Version:     1,
		ClickEvents: false,
	}

	hdrb, err := json.Marshal(hdr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hdrb))
	printStatus()
}
