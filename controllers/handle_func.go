package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

type User struct {
	Name    string `json:"name" form:"name"`
	Message string `json:"message" form:"message"`
	Nick    string `json:"nick" form:"nick"`
}

// application/x-www-form-urlencoded
func urlencodedPost(c *gin.Context) {

	name := c.Query("name")
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "1231412")

	log.Println(name, message, nick)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"name":    name,
		"message": message,
		"nick":    nick,
	})

}

// application/json  application/x-www-form-urlencoded
func jsonAndFormPost(c *gin.Context) {
	var user User

	c.Bind(&user)

	log.Println(user.Name, user.Message, user.Nick)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"name":    user.Name,
		"message": user.Message,
		"nick":    user.Nick,
	})

}

// 表单提交
func formPost(c *gin.Context) {

	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	log.Println(message, nick)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
		"nick":    nick,
	})
}

func xmlPost(c *gin.Context) {

}

func fileUpload(c *gin.Context) {

	filesUrl := make([]string, 0)

	form, err := c.MultipartForm()

	if err != nil {
		log.Println("postMultipleFile error: %s")
	}

	files := form.File["file"]

	_, err = os.Stat("upload")

	if err != nil {
		os.Mkdir("upload", os.ModePerm)
	}

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		if err = c.SaveUploadedFile(file, "upload/"+file.Filename); err != nil {
			log.Println("SaveUploadedFile error: %s")

			return
		}
		filesUrl = append(filesUrl, "upload/"+file.Filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"state": "SUCCESS",
		"url":   strings.Join(filesUrl, ";"),
	})
}
