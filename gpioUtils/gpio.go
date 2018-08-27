package gpioUtils

/**
* @author Mathieu GARREAU
* Aout 2018
 */

import (
	//"github.com/mlgd/gpio"
	"github.com/davecheney/gpio"
	"fmt"
	"time"
)

// list of gpios available on the raspberry pi
var gpios = []int {0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 17, 18, 22, 23, 24, 25}

func IsAvailableGpio(gpio int) bool {
	gpioMap := make(map[int]struct{}, len(gpios))
	for _, pin := range gpios {
		gpioMap[pin] = struct{}{}
	}

	_, ok := gpioMap[gpio]
	return ok
}

func Blink(gpioToOpen, nbBlink int) {
	// Open the port gpioToOpen with mode OUT
	pin, err := gpio.OpenPin(gpioToOpen, gpio.ModeOutput)
	if err != nil {
		fmt.Println("Error while opening pin: /n", err)
		return
	}

	// Do the DEL blink
	var i = 0
	for i < nbBlink {
		i++
		// switch ON
		pin.Set()
		// Wait
		time.Sleep(time.Second / 4 )
		// switch OFF
		pin.Clear()
		// wait
		time.Sleep(time.Second / 4)
	}
}

func SwitchOn(gpioToOpen int) bool {
	pin, err := gpio.OpenPin(gpioToOpen, gpio.ModeOutput)
	if err != nil {
		return false
	}
	pin.Set()
	return true
}

func SwitchOff(gpioToClose int) bool {
	pin, err := gpio.OpenPin(gpioToClose, gpio.ModeOutput)
	if err != nil {
		return false
	}
	pinStatus := pin.Get()
	if pinStatus == true {
		pin.Clear()
	}
	return true
}

func GPIOStatus(gpioForStatus int) (ok bool, err error) {
	pin, err := gpio.OpenPin(gpioForStatus, gpio.ModeOutput)
	if err != nil {
		return false, err
	}
	return pin.Get(), err
}
