package main

import (
	"net/http"
	"strconv"

	"github.com/alexcoder04/friendly/v2"
	"github.com/alexcoder04/legendum/processor"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ApiHandler(c *gin.Context) {
	starttimeRequest := c.Param("starttime")
	starttime := 0
	if friendly.IsInt(starttimeRequest) {
		starttime, _ = strconv.Atoi(starttimeRequest)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   processor.Load(starttime),
	})
}
