package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("./templates/*.html")

	r.GET("/", IndexHandler)
	r.GET("/load/:starttime", ApiHandler)

	r.Run(":6363")
}
