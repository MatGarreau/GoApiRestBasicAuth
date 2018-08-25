package main

/**
* @author Mathieu GARREAU
* Aout 2018
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"net/http"
	"apiREST/gpioUtils"
)

const (
	pinParamError = "An error occurred while getting pin param"
	gpioError     = "An error occurred while opening gpio"
)
func status() gin.HandlerFunc {
	return func(ctx *gin.Context){
		msg := "This server is up. Plz contact sys admin to use it!"
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
	}
}

func switchOn() gin.HandlerFunc {
	return func(ctx *gin.Context){
		user := ctx.Value(gin.AuthUserKey).(string)
		ok := gpioUtils.SwitchOn(17)
		if ok == false {
			ctx.JSON(http.StatusInternalServerError, gpioError)
			return
		}
		msg := fmt.Sprintf("switchon has been called by authenticated user: %s", user)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
		return
	}
}

func switchOff() gin.HandlerFunc {
	return func(ctx *gin.Context){
		// get user name from the context
		user := ctx.Value(gin.AuthUserKey).(string)
		ok := gpioUtils.SwitchOff(17)
		if ok == false {
			ctx.JSON(http.StatusInternalServerError, gpioError)
			return
		}
		msg := fmt.Sprintf("switchoff has been called by authenticated user: %s", user)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
		return
	}
}

func SwitchOnPin() gin.HandlerFunc {
	return func(ctx *gin.Context){
		user := ctx.Value(gin.AuthUserKey).(string)
		// get the pin number from the URL and convert it to int
		pin, err := strconv.Atoi(ctx.Param("pin"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pinParamError)
			return
		}
		if 0 < pin && pin < 25 {
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
		if 0 < pin && pin < 25 {
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
		return
	}
}

func main() {
	router := gin.New()

	// Router without authentication
	open := router.Group("/")
	open.GET("/status", status())


	// RouterGroup using gin.BasicAuth()
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":"bar",
	}))

	authorized.GET("/switchon", switchOn())

	authorized.GET("/switchoff", switchOff())

	authorized.GET("/switchonpin/:pin", SwitchOnPin())

	authorized.GET("/switchoffpin/:pin", SwitchOffPin())

	// Listen and serve on localhost:8088
	router.Run(":8088")
}