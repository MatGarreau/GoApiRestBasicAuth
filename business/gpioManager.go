package business

import (
	"fmt"
	"net/http"
	"strconv"
	"apiREST/gpioUtils"
	"github.com/gin-gonic/gin"
)

const (
	gpioParamError   = "An error occurred while getting pin param"
	gpioError        = "An error occurred while opening gpio"
	gpioNotAvailable = "The gpio number defined in URL is not available"
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

func getGPIONumber(ctx *gin.Context) (pin int, err error) {
	pin, err = strconv.Atoi(ctx.Param("gpio"))
	return pin, err
}

func GetGPIOStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := getUser(ctx)
		// get the gpio number from the URL and convert it to int
		gpio, err := getGPIONumber(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gpioParamError)
			return
		}
		// check the gpio is available for the raspberry pi
		if gpioUtils.IsAvailableGpio(gpio) == false {
			// gpio not available
			ctx.JSON(http.StatusBadRequest, gpioNotAvailable)
			return
		}
		// get gpio status (true == high level - false == low level)
		gpioStatus, err := gpioUtils.GPIOStatus(gpio)
		if err != nil {
			// error while getting gpio status
			ctx.JSON(http.StatusInternalServerError, gpioError)
			return
		}
		msg := fmt.Sprintf("gpio status %d has been called by authenticated user %s. GPIO status is: %t", gpio, user, gpioStatus)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg) // todo : return something like {"gpioID":gpio,"status":up} (or up/down) rather than a message
		return
	}
}

func SwitchOnGPIO() gin.HandlerFunc {
	return func(ctx *gin.Context){
		user := getUser(ctx)
		// get the gpio number from the URL and convert it to int
		gpio, err := getGPIONumber(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gpioParamError)
			return
		}
		// check the gpio is available for the raspberry pi
		if gpioUtils.IsAvailableGpio(gpio) == false {
			// gpio not available
			ctx.JSON(http.StatusBadRequest, gpioNotAvailable)
			return
		}
		// switch on gpio (set high level)
		ok := gpioUtils.SwitchOn(gpio)
		if ok == false {
			ctx.JSON(http.StatusInternalServerError, gpioError)
			return
		}
		msg := fmt.Sprintf("switchon gpio %d has been called by authenticated user: %s", gpio, user)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
		return
	}
}

func SwitchOffGPIO() gin.HandlerFunc {
	return func(ctx *gin.Context){
		user := getUser(ctx)
		// get the gpio number from the URL and convert it to int
		gpio, err := strconv.Atoi(ctx.Param("gpio"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gpioParamError)
			return
		}
		// check the gpio is available for the raspberry pi
		if gpioUtils.IsAvailableGpio(gpio) == false {
			// gpio not available
			ctx.JSON(http.StatusBadRequest, gpioNotAvailable)
			return
		}
		// switch off gpio (set low level)
		ok:= gpioUtils.SwitchOff(gpio)
		if ok == false {
			ctx.JSON(http.StatusInternalServerError, gpioError)
			return
		}
		msg := fmt.Sprintf("switchoff gpio %d has been called by authenticated user: %s", gpio, user)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
		return
	}
}

