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

func Blink(pinToOpen, nbBlink int) {
	// Open the port pinToOpen with mode OUT
	pin, err := gpio.OpenPin(pinToOpen, gpio.ModeOutput)
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

func SwitchOn(pinToOpen int) bool {
	pin, err := gpio.OpenPin(pinToOpen, gpio.ModeOutput)
	if err != nil {
		//fmt.Printf("Error while opening pin %s! %s\n", pin, err)
		return false
	}
	pin.Set()
	return true
}

func SwitchOff(pinToClose int) bool {
	pin, err := gpio.OpenPin(pinToClose, gpio.ModeOutput)
	if err != nil {
		//fmt.Printf("Error while opening pin %s! %s\n", pin, err)
		panic(err)
		return false
	}
	//pin.Clear()
	pinStatus := pin.Get()
	if pinStatus == true {
		pin.Clear()
	}
	return true
}

func PinStatus(pinForStatus int) (ok bool, err error) {
	pin, err := gpio.OpenPin(pinForStatus, gpio.ModeOutput)
	if err != nil {
		//fmt.Printf("Error while opening pin %s! %s\n", pin, err)
		panic(err)
		return false, err
	}
	return pin.Get(), err
}
