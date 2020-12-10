package controllers

import (
	"gin-router-web/serialize"
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
	nick := c.DefaultPostForm("nick", "default nick")
	name := c.DefaultPostForm("name", "default name")
	user := User{
		Name:    name,
		Nick:    nick,
		Message: message,
	}

	// This way is better
	// 下面这种方式 会自动和定义的结构体进行绑定
	// user := &User{}
	// c.ShouldBind(user)

	c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusOK, "success", user))
}

// UrlencodedPost application/x-www-form-urlencoded
func UrlencodedPost(c *gin.Context) {

	limit := c.Query("limit")
	name := c.PostForm("name")
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "1231412")
	user := User{
		Name:    name,
		Nick:    nick,
		Message: message,
	}

	// This way is better
	// 下面这种方式 会自动和定义的结构体进行绑定
	// user := &User{}
	// c.ShouldBind(user)

	log.Printf("request query limit: %s\n", limit)

	c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusOK, "success", user))
}

// JSONPost json
func JSONPost(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusBadRequest, "fail", nil))
		return
	}

	c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusOK, "success", user))
}

//JSONAndFormPost  application/json  application/x-www-form-urlencoded
func JSONAndFormPost(c *gin.Context) {
	var user User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusBadRequest, "fail", nil))
		return
	}

	c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusOK, "success", user))
}

//XMLPost xml
func XMLPost(c *gin.Context) {
	var user User

	// c.ShouldBind(&user)
	// c.Bind(&user)
	if err := c.BindXML(&user); err != nil {
		c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusBadRequest, "fail", nil))
		return
	}

	c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusOK, "success", user))
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

	c.JSON(http.StatusOK, serialize.BuildResponse(http.StatusOK, "success", gin.H{
		"url": strings.Join(filesURL, ";"),
	}))
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

	// http://c.biancheng.net/view/5729.html
	// 打开文件并追加内容
	// O_RDWR：读写模式打开文件, O_CREATE：如果不存在将创建一个新文件, O_APPEND：写操作时将数据附加到文件尾部（追加）
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
