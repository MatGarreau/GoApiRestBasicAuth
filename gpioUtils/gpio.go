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

func SwitchOn(pinToOpen int){
	pin, err := gpio.OpenPin(pinToOpen, gpio.ModeOutput)
	if err != nil {
		fmt.Printf("Error while opening pin %s! %s\n", pin, err)
		return
	}
	pin.Set()
}

func SwitchOff(pinToClose int){
	pin, err := gpio.OpenPin(pinToClose, gpio.ModeOutput)
	if err != nil {
		fmt.Printf("Error while opening pin %s! %s\n", pin, err)
		return
	}
	pin.Clear()
}