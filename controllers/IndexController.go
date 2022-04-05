package controllers

import (
	"github.com/gin-gonic/gin"
)

func IndexHomePage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Homepage",
	})
	return
}
