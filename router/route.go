package router

import (
	"net/http"

	"gin-router-web/api"
	"gin-router-web/models"

	"github.com/gin-gonic/gin"
)

func setStaticFS(r *gin.Engine) {
	// set html template
	r.LoadHTMLGlob("./views/*.html")

	// set server static
	r.StaticFile("favicon.ico", "./public/favicon.ico")
	r.StaticFS("/static", http.Dir("public/static"))
	r.StaticFS("/upload", http.Dir("upload"))
}

func setWebRoute(r *gin.Engine) {
	// 首页 router /
	r.GET("/", api.WebIndex)
	r.GET("/upload_chunks", api.WebUploadChunks)
}

// SetupRouter  set gin router
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// 设置静态资源
	setStaticFS(r)

	// set web router
	setWebRoute(r)

	apiGroup := r.Group("/api")
	{
		// 表单提交
		apiGroup.POST("/form_post", api.FormPost)

		// json提交
		apiGroup.POST("/json_post", api.JSONPost)

		//url encode 提交
		apiGroup.POST("/urlencoded_post", api.UrlencodedPost)

		// 即支持json又支持form
		apiGroup.POST("/json_and_form_post", api.JSONAndFormPost)

		// xml 提交
		apiGroup.POST("/xml_post", api.XMLPost)

		// 文件上传
		apiGroup.POST("/file_upload", api.FileUpload)

		// 文件分片上传
		apiGroup.POST("/file_chunk_upload", api.FileChunkUpload)

		apiGroup.GET("/query", func(c *gin.Context) {
			message := c.Query("message")
			nick := c.DefaultQuery("nick", "anonymous")

			c.JSON(http.StatusOK, models.BuildResponse(http.StatusOK, "success", gin.H{
				message: message,
				nick:    nick,
			}))
		})
	}

	return r
}
