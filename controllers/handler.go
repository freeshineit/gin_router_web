package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setStaticFS(r *gin.Engine) {
	r.LoadHTMLGlob("views/*")

	r.StaticFile("favicon.ico", "./views/favicon.ico")
	r.StaticFS("/static", http.Dir("public/static"))
	r.StaticFS("/upload", http.Dir("upload"))
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	setStaticFS(r) // 设置静态资源

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	api := r.Group("/api")
	{
		api.POST("/form_post", formPost)

		api.POST("/json_post", jsonPost)
		api.POST("/urlencoded_post", urlencodedPost)
		api.POST("/json_and_form_post", jsonAndFormPost)
		api.POST("/xml_post", xmlPost)
		api.POST("/file_upload", fileUpload)

		api.POST("/file_chunk_upload", fileChunkUpload)

		api.GET("/list", func(c *gin.Context) {
			message := c.Query("message")
			nick := c.DefaultQuery("nick", "anonymous")

			c.JSON(http.StatusOK, gin.H{
				"status":  "SUCCESS",
				"message": message,
				"nick":    nick,
			})
		})
	}

	return r
}
