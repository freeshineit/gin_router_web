package router

import (
	"net/http"

	"gin-router-web/controllers"
	"gin-router-web/serialize"

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

func setWebRouter(r *gin.Engine) {
	// 首页 router /
	r.GET("/", controllers.WebIndex)
	r.GET("/upload", controllers.WebUpload)
}

// SetupRouter  set gin router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 设置静态资源
	setStaticFS(r)

	// set web router
	setWebRouter(r)

	api := r.Group("/api")
	{
		// 表单提交
		api.POST("/form_post", controllers.FormPost)

		// json提交
		api.POST("/json_post", controllers.JSONPost)

		//url encode 提交
		api.POST("/urlencoded_post", controllers.UrlencodedPost)

		// 即支持json又支持form
		api.POST("/json_and_form_post", controllers.JSONAndFormPost)

		// xml 提交
		api.POST("/xml_post", controllers.XMLPost)

		// 文件上传
		api.POST("/file_upload", controllers.FileUpload)

		// 文件分片上传
		api.POST("/file_chunk_upload", controllers.FileChunkUpload)

		api.GET("/query", func(c *gin.Context) {
			message := c.Query("message")
			nick := c.DefaultQuery("nick", "anonymous")

			c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusOK, "success", gin.H{
				message: message,
				nick:    nick,
			}))
		})
	}

	return r
}
