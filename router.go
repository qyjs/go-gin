package main

import (
	"gin/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", apis.IndexAPI)

	router.POST("/person", apis.AddPersonAPI) // done

	router.GET("/person/:id", apis.GetPersonsAPI) // done

	router.GET("/person", apis.GetPersonAPI) // done

	router.PUT("/person/:id", apis.ModPersonAPI) // done

	router.DELETE("/person/:id", apis.DelPersonAPI)

	return router
}
