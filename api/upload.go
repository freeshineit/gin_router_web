package api

import (
	"gin-router-web/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// FileUpload file upload
// /api/file_upload
func FileUpload(c *gin.Context) {

	filesURL := make([]string, 0)

	form, err := c.MultipartForm()

	log.Println(c.Cookie("token"))
	log.Println(c.GetHeader("Content-Type"))

	if err != nil {
		log.Printf("postMultipleFile error: %s \n", err.Error())
		return
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

	c.JSON(http.StatusOK, models.BuildResponse(http.StatusOK, "success", gin.H{
		"urls": filesURL,
	}))
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

	var chunkFile models.ChunkFile
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
		c.JSON(http.StatusOK,
			models.BuildResponse(http.StatusOK, "success", gin.H{"url": "/" + filePath}),
		)
	} else {
		contentType := strings.Split(c.GetHeader("Content-Type"), "boundary=")
		c.String(http.StatusOK, contentType[1])
	}
}
