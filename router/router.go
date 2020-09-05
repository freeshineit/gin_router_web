package router

import (
	"net/http"

	"gin-router-web/controllers"

	"github.com/gin-gonic/gin"
)

func setStaticFS(r *gin.Engine) {
	// set html template
	r.LoadHTMLGlob("views/*")

	// set server static
	r.StaticFile("favicon.ico", "./views/favicon.ico")
	r.StaticFS("/static", http.Dir("public/static"))
	r.StaticFS("/upload", http.Dir("upload"))
}

// SetupRouter  set gin router
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
		api.POST("/form_post", controllers.FormPost)

		api.POST("/json_post", controllers.JSONPost)
		api.POST("/urlencoded_post", controllers.UrlencodedPost)
		api.POST("/json_and_form_post", controllers.JSONAndFormPost)
		api.POST("/xml_post", controllers.XMLPost)
		api.POST("/file_upload", controllers.FileUpload)

		api.POST("/file_chunk_upload", controllers.FileChunkUpload)

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
