package main

/**
* @author Mathieu GARREAU
* Aout 2018
 */

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"apiREST/gpioUtils"
	"net/http"
)

func status() gin.HandlerFunc {
	return func(ctx *gin.Context){
		fmt.Println("This server is up. Plz contact sys admin to use it!")
	}
}

func switchOn() gin.HandlerFunc {
	return func(ctx *gin.Context){
		gpioUtils.SwitchOn(17)
		fmt.Println("switch ON ")
		user := ctx.MustGet(gin.AuthUserKey).(string)
		fmt.Println("switchon has been called by:", user)
		return
	}
}

func switchOff() gin.HandlerFunc {
	return func(ctx *gin.Context){
		gpioUtils.SwitchOff(17)
		fmt.Println("switch OFF ")
		user := ctx.MustGet(gin.AuthUserKey).(string)
		msg := fmt.Sprintf("switchoff has been called by: %s", user)
		fmt.Println(msg)
		ctx.JSON(http.StatusOK, msg)
		return
	}
}

func SwitchOnPin() gin.HandlerFunc {
	return func(ctx *gin.Context){
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

func main() {
	router := gin.New()

	// Router without authentication
	open := router.Group("/")
	open.GET("/status", status())


	// RouterGroup using gin.BasicAuth()
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
	}))

	authorized.POST("/switchon", switchOn())

	authorized.POST("/switchoff", switchOff())

	authorized.POST("/switchonpin/{pin}", SwitchOnPin())

	// Listen and serve on localhost:8088
	router.Run(":8088")
}