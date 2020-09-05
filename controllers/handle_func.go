package controllers

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// User user struct
type User struct {
	Name    string `json:"name" form:"name" xml:"name"`
	Message string `json:"message" form:"message" xml:"message"`
	Nick    string `json:"nick" form:"nick" xml:"nick"`
}

// FormPost 表单提交
func FormPost(c *gin.Context) {

	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	log.Println(message, nick)
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": message,
		"nick":    nick,
	})
}

// UrlencodedPost application/x-www-form-urlencoded
func UrlencodedPost(c *gin.Context) {

	name := c.Query("name")
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "1231412")

	log.Println(name, message, nick)
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"name":    name,
		"message": message,
		"nick":    nick,
	})
}

// JSONPost json
func JSONPost(c *gin.Context) {
	var user User

	c.BindJSON(&user)

	log.Println(user.Name, user.Message, user.Nick)

	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"name":    user.Name,
		"message": user.Message,
		"nick":    user.Nick,
	})
}

//JSONAndFormPost  application/json  application/x-www-form-urlencoded
func JSONAndFormPost(c *gin.Context) {
	var user User

	c.Bind(&user)

	log.Println(user.Name, user.Message, user.Nick)

	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"name":    user.Name,
		"message": user.Message,
		"nick":    user.Nick,
	})
}

//XMLPost xml
func XMLPost(c *gin.Context) {
	var user User

	c.Bind(&user)

	log.Println(user.Name, user.Message, user.Nick)

	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"name":    user.Name,
		"message": user.Message,
		"nick":    user.Nick,
	})
}

// FileUpload file upload
func FileUpload(c *gin.Context) {

	filesURL := make([]string, 0)

	form, err := c.MultipartForm()

	log.Println(c.Cookie("token"))
	log.Println(c.GetHeader("Content-Type"))

	if err != nil {
		log.Printf("postMultipleFile error: %s \n", err.Error())
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
			log.Printf("SaveUploadedFile error: %s \n", err.Error())

			return
		}
		filesURL = append(filesURL, "upload/"+file.Filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"state": "SUCCESS",
		"url":   strings.Join(filesURL, ";"),
	})
}

// ChunkFile file chunk
type ChunkFile struct {
	Name   string                `json:"name" form:"name"`
	Chunk  int                   `json:"chunk" form:"chunk"`
	Chunks int                   `json:"chunks" form:"chunks"`
	File   *multipart.FileHeader `json:"file" form:"file"`
}

// PathExists 判断文件是否已经存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// FileChunkUpload file chunk upload
func FileChunkUpload(c *gin.Context) {

	var chunkFile ChunkFile
	r := c.Request

	c.Bind(&chunkFile)

	var Buf = make([]byte, 0)
	// in your case file would be fileupload
	file, _, _ := r.FormFile("file")

	//log.Println(reflect.TypeOf(chunkFile.File))
	//log.Println(reflect.TypeOf(file))
	log.Println("this is ", chunkFile.File)
	Buf, _ = ioutil.ReadAll(file)

	filePath := "upload/" + chunkFile.Name

	bool, _ := PathExists(filePath)

	if !bool {
		os.Create(filePath)
	}
	fd, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.Write(Buf)
	fd.Close()

	if chunkFile.Chunk+1 == chunkFile.Chunks {
		c.JSON(http.StatusOK, gin.H{
			"state": "SUCCESS",
			"url":   "/" + filePath,
		})
	} else {
		contentType := strings.Split(c.GetHeader("Content-Type"), "boundary=")
		c.String(http.StatusOK, contentType[1])
	}
}
