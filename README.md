# DMX for Go

A simple Go package to send DMX messages.

Copyright (c) 2013 AKUALAB INC. All Rights Reserved.
http://www.akualab.com - @akualab - info@akualab.com

## Credits
Ported from [pySimpleDMX](https://github.com/c0z3n/pySimpleDMX) by @c0z3n

## DMX Hardware
To send DMX messages using a serial port over USB, you can use one of these adaptors:
* Enttec [DMX USB Pro](http://www.enttec.com/?main_menu=Products&pn=70304)
* DMX King [UltraDMX Micro](http://dmxking.com/usbdmx/ultradmxmicro)

You can buy an 86 RGB LED Light PAR DMX-512 on eBay for $30.

## Install
go get github.com/akualab/dmx

## Example

```
package main

import (
    "github.com/akualab/dmx"
)

func main() {

    dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
    if e != nil {
        log.Fatal(e)
    }

    // Set values for channels.
    dmx.SetChannel(1, 100)
    dmx.SetChannel(2, 70)
    dmx.SetChannel(3, 130)
    dmx.SetChannel(4, 180)

    // Send!
    dmx.Render()
}
```

## Run Random Color Example
go run example/random-color.go

## License
GPL Version 3 - http://www.gnu.org/licenses/gpl.html

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.
