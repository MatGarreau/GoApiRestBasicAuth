package main

/**
* @author Mathieu GARREAU
* Aout 2018
 */

import (
	"github.com/gin-gonic/gin"
	"apiREST/business"
)




func main() {
	router := gin.New()

	// Router without authentication
	open := router.Group("/")
	open.GET("/status", business.Status())


	// RouterGroup using gin.BasicAuth()
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":"bar",
	}))

	authorized.GET("/gpiostatus/:gpio", business.GetGPIOStatus())

	authorized.PUT("/switchongpio/:gpio", business.SwitchOnGPIO())

	authorized.PUT("/switchoffgpio/:gpio", business.SwitchOffGPIO())

	// Listen and serve on localhost:8088
	router.Run(":8088")
}