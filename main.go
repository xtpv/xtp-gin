package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := HttpStart()
	if err != nil {
		log.Panic(err)
	}
}

func HttpStart() error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "success"})
	})
	return r.Run()
}