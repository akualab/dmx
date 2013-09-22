package main

import (
	"github.com/akualab/dmx"
	"log"
	"math/rand"
	"time"
)

func main() {

	r := rand.New(rand.NewSource(99))
	log.Printf("start DMX")
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		log.Fatal(e)
	}

	// Send RGB
	// 1: brightness/flash, 2: red, 3: blue, 4: green
	dmx.ChannelMap(1, 2, 4, 3)
	dmx.SendRGB(130, 100, 100, 100)
	time.Sleep(10 * time.Second)

	// Initial color.
	dmx.SetChannel(1, 100)
	dmx.SetChannel(2, 70)
	dmx.SetChannel(3, 130)
	dmx.SetChannel(4, 180)
	dmx.Render()

	for {

		// Wait.
		time.Sleep(100 * time.Millisecond)

		dmx.SetChannel(1, 20)                // Intensity
		dmx.SetChannel(2, byte(r.Intn(256))) // R
		dmx.SetChannel(3, byte(r.Intn(256))) // G
		dmx.SetChannel(4, byte(r.Intn(256))) // B
		dmx.Render()

	}
}
