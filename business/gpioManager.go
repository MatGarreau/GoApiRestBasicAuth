package business

import (
	"fmt"
	"net/http"
	"strconv"
	"apiREST/gpioUtils"
	"github.com/gin-gonic/gin"
)

const (
	pinParamError = "An error occurred while getting pin param"
	gpioError     = "An error occurred while opening gpio"
	pinNotAvailable = "The pin number defined in URL is not available"
)

func Status() gin.HandlerFunc {
	return func(ctx *gin.Context){
		msg := "This server is up. Plz contact sys admin to use it!"
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
	}
}

func getUser(ctx *gin.Context) string {
	return ctx.Value(gin.AuthUserKey).(string)
}

func getPinNumber(ctx *gin.Context) (pin int, err error) {
	pin, err = strconv.Atoi(ctx.Param("pin"))
	return pin, err
}

func GetPinStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := getUser(ctx)
		// get the pin number from the URL and convert it to int
		pin, err := getPinNumber(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pinParamError)
			return
		}
		// check the pin is available for the raspberry pi
		if gpioUtils.IsAvailableGpio(pin) == false {
			// pin not available
			ctx.JSON(http.StatusBadRequest, pinNotAvailable)
			return
		}
		// get pin status (true == high level - false == low level)
		pinStatus, err := gpioUtils.PinStatus(pin)
		if err != nil {
			// error while getting pin status
			ctx.JSON(http.StatusInternalServerError, gpioError)
			return
		}
		msg := fmt.Sprintf("pin status %d has been called by authenticated user %s. Pin status is: %t", pin, user, pinStatus)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg) // todo : return something like {"pinID":pinNumber,"status":up} (or up/down) rather than a message
		return
	}
}

func SwitchOnPin() gin.HandlerFunc {
	return func(ctx *gin.Context){
		user := getUser(ctx)
		// get the pin number from the URL and convert it to int
		pin, err := getPinNumber(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pinParamError)
			return
		}
		// check the pin is available for the raspberry pi
		if gpioUtils.IsAvailableGpio(pin) == false {
			// pin not available
			ctx.JSON(http.StatusBadRequest, pinNotAvailable)
			return
		}
		// switch on pin (set high level)
		ok := gpioUtils.SwitchOn(pin)
		if ok == false {
			ctx.JSON(http.StatusInternalServerError, gpioError)
			return
		}
		msg := fmt.Sprintf("switchon pin %d has been called by authenticated user: %s", pin, user)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
		return
	}
}

func SwitchOffPin() gin.HandlerFunc {
	return func(ctx *gin.Context){
		user := ctx.Value(gin.AuthUserKey).(string)
		// get the pin number from the URL and convert it to int
		pin, err := strconv.Atoi(ctx.Param("pin"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pinParamError)
			return
		}
		// check the pin is available for the raspberry pi
		if gpioUtils.IsAvailableGpio(pin) == false {
			// pin not available
			ctx.JSON(http.StatusBadRequest, pinNotAvailable)
			return
		}
		// switch off pin (set low level)
		ok:= gpioUtils.SwitchOff(pin)
		if ok == false {
			ctx.JSON(http.StatusInternalServerError, gpioError)
			return
		}
		msg := fmt.Sprintf("switchoff pin %d has been called by authenticated user: %s", pin, user)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
		return
	}
}

