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
	"github.com/davecheney/gpio"

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
			ctx.JSON(http.StatusInternalServerError, "An error occurred while opening gpio")
			return
		}
		msg := fmt.Sprintf("switchon has been called by: %s", user)
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
			ctx.JSON(http.StatusInternalServerError, "An error occurred while opening gpio")
			return
		}
		msg := fmt.Sprintf("switchoff has been called by: %s", user)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
		return
	}
}

func SwitchOnPin() gin.HandlerFunc {
	return func(ctx *gin.Context){
		// get the pin number from the URL and convert it to int
		pin, err := strconv.Atoi(ctx.Param("pin"))
		if err != nil {
			panic(err)
		}
		if 0 < pin && pin < 25 {
			gpioUtils.SwitchOn(pin)
		}
		return
	}
}


func SwitchOffPin() gin.HandlerFunc {
	return func(ctx *gin.Context){
		// get the pin number from the URL and convert it to int
		pin, err := strconv.Atoi(ctx.Param("pin"))
		if err != nil {
			panic(err)
		}
		if 0 < pin && pin < 25 {
			gpioUtils.SwitchOff(pin)
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
		"foo":    "bar",
	}))

	authorized.GET("/switchon", switchOn())

	authorized.GET("/switchoff", switchOff())

	authorized.GET("/switchonpin/{pin}", SwitchOnPin())

	authorized.GET("/switchoffpin/{pin}", SwitchOffPin())

	// Listen and serve on localhost:8088
	router.Run(":8088")
}