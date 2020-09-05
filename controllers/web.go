package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WebIndex  index web.
// router [/]
func WebIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "this is title",
	})
}

// WebUpload upload web
// router [/upload.html]
func WebUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}
