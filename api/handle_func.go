package api

import (
	"gin-router-web/helper"
	"gin-router-web/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FormPost 表单提交
func FormPost(c *gin.Context) {

	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "default nick")
	name := c.DefaultPostForm("name", "default name")
	user := models.User{
		Name:    name,
		Nick:    nick,
		Message: message,
	}

	// This way is better
	// 下面这种方式 会自动和定义的结构体进行绑定
	// user := &User{}
	// c.ShouldBind(user)

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}

// UrlencodedPost application/x-www-form-urlencoded
func UrlencodedPost(c *gin.Context) {

	limit := c.Query("limit")
	name := c.PostForm("name")
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "1231412")
	user := models.User{
		Name:    name,
		Nick:    nick,
		Message: message,
	}

	// This way is better
	// 下面这种方式 会自动和定义的结构体进行绑定
	// user := &User{}
	// c.ShouldBind(user)

	log.Printf("request query limit: %s\n", limit)

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}

// JSONPost json
func JSONPost(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse[any](http.StatusBadRequest, "invalid parameter"))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}

// JSONAndFormPost  application/json  application/x-www-form-urlencoded
func JSONAndFormPost(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse[any](http.StatusBadRequest, "invalid parameter"))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}

// XMLPost xml
func XMLPost(c *gin.Context) {
	var user models.User

	// c.ShouldBind(&user)
	// c.Bind(&user)
	if err := c.BindXML(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse[any](http.StatusBadRequest, "invalid parameter"))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
