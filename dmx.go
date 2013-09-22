// Simple Go package to send DMX messages.
// Copyright (c) 2013 AKUALAB INC. All Rights Reserved.
// www.akualab.com - @akualab - info@akualab.com
//
// CREDITS:
// Ported from pySimpleDMX (https://github.com/c0z3n/pySimpleDMX)
// Written by Michael Dvorkin
//
// GNU General Public License v3.  http://www.gnu.org/licenses/
package dmx

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"log"
)

const (
	START_VAL       = 0x7E
	END_VAL         = 0xE7
	BAUD            = 57600
	TIMEOUT         = 1
	DEV             = "/dev/ttyUSB0"
	FRAME_SIZE      = 511
	FRAME_SIZE_LOW  = byte(FRAME_SIZE & 0xFF)
	FRAME_SIZE_HIGH = byte(FRAME_SIZE >> 8 & 0xFF)
)

var labels = map[string]byte{
	"GET_WIDGET_PARAMETERS": 3, // unused
	"SET_WIDGET_PARAMETERS": 4, // unused
	"RX_DMX_PACKET":         5, // unused
	"TX_DMX_PACKET":         6,
	"TX_RDM_PACKET_REQUEST": 7, // unused
	"RX_DMX_ON_CHANGE":      8, // unused
}

// A serial DMX connection.
type DMX struct {
	dev            string
	frame          [FRAME_SIZE]byte
	packet         [FRAME_SIZE + 10]byte
	serial         io.ReadWriteCloser
	redChan        int
	blueChan       int
	greenChan      int
	brightnessChan int
}

// Creates a new DMX connection using a serial device.
func NewDMXConnection(device string) (dmx *DMX, err error) {

	dmx = &DMX{}

	// Set serial device or use default.
	dmx.dev = device
	if len(dmx.dev) == 0 {
		dmx.dev = DEV
	}

	c := &serial.Config{Name: dmx.dev, Baud: BAUD}
	dmx.serial, err = serial.OpenPort(c)
	if err != nil {
		return
	}
	log.Printf("Opened port [%s].", dmx.dev)
	return
}

// Set channel level in the dmx frame to be rendered
// the next time Render() is called.
func (dmx *DMX) SetChannel(channel int, val byte) error {

	checkChannelID(channel)
	dmx.frame[channel] = val
	return nil
}

// Turn off a specific channel.
func (dmx *DMX) ClearChannel(channel int) error {

	checkChannelID(channel)
	dmx.frame[channel] = 0
	return nil
}

// Turn off all channels.
func (dmx *DMX) ClearAll() {

	for i, _ := range dmx.frame {
		dmx.frame[i] = 0
	}
}

// Send frame to serial device.
func (dmx *DMX) Render() error {

	p := dmx.packet[:0]
	p = append(p, START_VAL)
	p = append(p, labels["TX_DMX_PACKET"])
	p = append(p, FRAME_SIZE_LOW)
	p = append(p, FRAME_SIZE_HIGH)
	p = append(p, dmx.frame[0:]...)
	p = append(p, END_VAL)

	// Write dmx frame.
	_, err := dmx.serial.Write(p)
	if err != nil {
		return err
	}
	return nil
}

// Close serial port.
func (dmx *DMX) Close() error {
	return dmx.serial.Close()
}

// Convenience method to map colors and brightness to channels.
func (dmx *DMX) ChannelMap(brightness, red, green, blue int) {

	checkChannelID(brightness)
	checkChannelID(red)
	checkChannelID(green)
	checkChannelID(blue)

	dmx.brightnessChan = brightness
	dmx.redChan = red
	dmx.greenChan = green
	dmx.blueChan = blue
}

// Configures RGB+Brightness channels and renders the color.
// Call ChannelMap to configure the RGB channels before calling
// this method.
func (dmx *DMX) SendRGB(brightness, red, green, blue byte) (e error) {

	dmx.ClearAll()
	e = dmx.SetChannel(dmx.brightnessChan, brightness)
	if e != nil {
		return
	}
	e = dmx.SetChannel(dmx.redChan, red)
	if e != nil {
		return
	}
	e = dmx.SetChannel(dmx.greenChan, green)
	if e != nil {
		return
	}
	e = dmx.SetChannel(dmx.blueChan, blue)
	if e != nil {
		return
	}
	e = dmx.Render()
	if e != nil {
		return
	}
	return
}

func checkChannelID(id int) {
	if (id > 512) || (id < 1) {
		panic(fmt.Sprintf("Invalid channel [%d]", id))
	}
}
