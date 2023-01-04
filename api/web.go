package api

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

// WebUploadChunks upload web
// router [/upload.html]
func WebUploadChunks(c *gin.Context) {
	c.HTML(http.StatusOK, "upload_chunks.html", gin.H{})
}
