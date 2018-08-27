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

	authorized.GET("/pinstatus/:pin", business.GetPinStatus())

	authorized.PUT("/switchonpin/:pin", business.SwitchOnPin())

	authorized.PUT("/switchoffpin/:pin", business.SwitchOffPin())

	// Listen and serve on localhost:8088
	router.Run(":8088")
}