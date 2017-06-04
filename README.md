# epaper

## Overview

This is a Go library that exposes the interface of Waveshare's
[4.3inch e-Paper UART Module](http://www.waveshare.com/wiki/4.3inch_e-Paper).

Most of the interface follows the wiki page above, but some of it was derived
by experimenting with the device.

## Usage

### Connecting

To connect the e-Paper module, you'll probably need a *USB to TTL serial* cable.
These are cheap and can be found in many online stores that sell Arduino and
Raspberry Pi accessories. Once you connect it to your PC, you need to find out
which port it uses. On Linux, it'll likely be `/dev/ttyUSB0`. On Windows, use
Device Manager to find out.

If you intend to use it on the Raspberry Pi, you can use the Pi's own serial
interface. See [here](http://elinux.org/RPi_Serial_Connection#Connection_to_a_microcontroller_or_other_peripheral)
for more info - specifically, on how to prevent Linux from using the serial
port. The serial port on a Raspberry Pi is either `/dev/ttyAMA0` (prior to
Raspberry Pi 3) or `/dev/ttyS0` (Raspberry Pi 3).

### Drawing things

Drawing on the device is achieved by first sending one or more commands which
modify the "screen state" on the device itself and then calling the `Update`
command which draws the current state.

The following commands are able to modify the "screen state":

-   Manipulating a single pixel
    -   `FillPixel`
    -   `ColorPixel`
-   Drawing shapes
    -   `DrawLine`
    -   `ColorLine`
    -   `DrawRect`
    -   `DrawCircle`
    -   `DrawTriangle`
-   Filling shapes
    -   `FillRect`
    -   `FillCircle`
    -   `FillTriangle`
-   Whole screen manipulation
    -   `Clear`
    -   `FillScreen`
- Miscellaneous
    -   `DisplayText`
    -   `DisplayImage`

### Example

```go
  ep, err := epaper.New("/dev/ttyUSB0")
  // handle error
  err = ep.DrawRect(100, 100, 200, 200)
  // handle error
  err = ep.Update()
  // handle error
  ep.Close()
```

For a more comprehensive example, see [example_test.go](https://github.com/fstanis/epaper/blob/master/example_test.go).

## Documentation

Please see the [GoDoc documentation](https://godoc.org/github.com/fstanis/epaper)
for documentation.

The [Waveshare wiki page](http://www.waveshare.com/wiki/4.3inch_e-Paper) also
provides useful information on what many of the commands do and how they are
used.
