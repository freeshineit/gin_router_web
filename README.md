# gin_router_web

![build](https://github.com/freeshineit/gin_rotuer_web/workflows/build/badge.svg)

[gin](https://github.com/gin-gonic/gin)是简单快速的`golang`框架,这篇文章主要是介绍`gin`的路由配置及使用（主要是post方法）

golang >= 1.18



## 使用

```bash
# development
go run main.go

# run development
# https://github.com/cosmtrek/air Live reload for Go apps
air

# build
go build
# or
make build

# run production
# export GIN_MODE=release
./gin-router-web

# server 8080
http://localhost:8080/
# file chunk upload
http://localhost:8080/upload_chunks

# docker deploy
make serve

##
go mod tidy
```

## API
```md
[POST]   /api/form_post          
[POST]   /api/json_post          
[POST]   /api/urlencoded_post    
[POST]   /api/json_and_form_post 
[POST]   /api/xml_post           
[POST]   /api/file_upload        
[POST]   /api/file_chunk_upload  
[GET]    /api/query              
```

## 静态资源配置

```go
func setStaticFS(r *gin.Engine) {
	// set html template
	r.LoadHTMLGlob("views/*")

	// set server static
	r.StaticFile("favicon.ico", "./public/favicon.ico")
	r.StaticFS("/static", http.Dir("public/static"))
	r.StaticFS("/upload", http.Dir("upload"))
}
```

`func (engine *Engine) LoadHTMLGlob(pattern string)`函数加载全局模式的 HTML 文件标识，并将结果与 HTML 渲染器相关联。

`func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes` 设置相对路径的静态资源

## api

> api 路由分组

```go
api := r.Group("/api")
{
	api.POST("/form_post", formPost)

	api.POST("/json_post", jsonPost)
	api.POST("/urlencoded_post", urlencodedPost)
	api.POST("/json_and_form_post", jsonAndFormPost)
	api.POST("/xml_post", xmlPost)
	api.POST("/file_upload", fileUpload)

	api.GET("/list", func(c *gin.Context) {
		message := c.Query("message")
		nick := c.DefaultQuery("nick", "anonymous")

    c.JSON(http.StatusOK, helper.BuildResponse(gin.H{
      message: message,
      nick:    nick,
    }))
	})
}

```

## 消息的类型

常用请求`Headers`中`Content-Type`的类型有`text/plain`、`text/html`、`application/json`、`application/x-www-form-urlencoded`、`application/xml`和`multipart/form-data`等.

- `text/plain` 纯文本
- `text/html` HTML 文档
- `application/json` json 格式数据
- `application/x-www-form-urlencoded` 使用 HTTP 的 POST 方法提交的表单
- `application/xml` xml 格式数据
- `application/form-data`主要是用来上传文件

[MIME](https://zh.wikipedia.org/wiki/MIME)

### form 表单提交

gin 路由实现

```go
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

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

html 实现

```html
<form method="post" action="/api/form_post" id="form">
  <div class="form-item">
    <label for="name">name</label>
    <input type="text" id="name" name="name" />
  </div>
  <div class="form-item">
    <label for="message">message</label>
    <input type="text" id="message" name="message" />
  </div>
  <div class="form-item">
    <label for="name">nick</label>
    <input type="text" id="nick" name="nick" />
  </div>
  <button type="submit">提交</button>
</form>
```

## post 提交`application/json`类型数据

gin 路由实现

```go
// JSONPost json
func JSONPost(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(http.StatusBadRequest, "invalid parameter"))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

js 实现

```js
axios({
  method: "post",
  url: "/api/json_post",
  headers: {
    "Content-Type": "application/json",
  },
  data,
}).then((res) => {
  console.log(res.data);
  $(".json-msg").text(`success  ${new Date()}`);
});
```

## post 提交`application/x-www-form-urlencoded`类型数据

gin 实现

```go
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

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

js 实现

```js
axios({
  method: "post",
  url: "/api/urlencoded_post?name=shineshao",
  headers: {
    "Content-Type": "application/x-www-form-urlencoded",
  },
  data: $.param(data),
}).then((res) => {
  console.log(res.data);
  $(".urlencoded-msg").text(`success  ${new Date()}`);
});
```

## post 提交`application/x-www-form-urlencoded`或`application/json`类型数据

gin

```go
//JSONAndFormPost  application/json  application/x-www-form-urlencoded
func JSONAndFormPost(c *gin.Context) {
	var user User

	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(http.StatusBadRequest, "invalid parameter"))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

js 实现

```js
// json
axios({
  method: "post",
  url: "/api/json_and_form_post",
  headers: {
    "Content-Type": "application/json",
  },
  data,
}).then((res) => {
  console.log(res.data);
  $(".jsonandform-msg").text(`success application/json data,  ${new Date()}`);
});

// x-www-form-urlencoded
axios({
  method: "post",
  url: "/api/json_and_form_post",
  headers: {
    "Content-Type": "application/x-www-form-urlencoded",
  },
  data: $.param(data),
}).then((res) => {
  console.log(res.data);
  $(".jsonandform-msg").text(
    `success application/x-www-form-urlencoded data${new Date()}`
  );
});
```

## post 提交`application/xml`类型数据(`application/xml`)

gin 实现

```go
//XMLPost xml
func XMLPost(c *gin.Context) {
	var user User

	// c.ShouldBind(&user)
	// c.Bind(&user)
	if err := c.BindXML(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(http.StatusBadRequest, "invalid parameter"))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

js 实现

```js
axios({
  method: "post",
  url: "/api/xml_post",
  headers: {
    "Content-Type": "application/xml",
  },
  data: `<xml><name>${data.name}</name><message>${data.message}</message><nick>${data.nick}</nick></xml>`,
});
```

## post 提交`multipart/form-data`类型数据(`multipart/form-data`)

gin 实现文件上传 (api/upload.go)

```go
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

	c.JSON(http.StatusOK, models.BuildResponse(gin.H{
		"urls": filesURL,
	}))
}
```

html 实现

```html
<div>
  <form id="multipleForm">
    <input
      type="file"
      name="file"
      id="file"
      multiple="multiple"
      accept="image/*"
    />
  </form>
  <button class="file_upload">开始上传文件</button>
</div>
```

js 实现

```js
// 单个文件上传
// var fd = new FormData()
// var file = document.getElementById('file')
// fd.append('file', file.files[0])

axios({
  method: "post",
  url: "/api/file_upload",
  headers: {
    "Content-Type": "application/form-data",
  },
  // data:fd // 单个文件上传
  data: new FormData($("#multipleForm")[0]),
}).then((res) => {
    console.log(res.data);
    const urls = res.data.data.urls || [];

    let imgHtml = "";

    for (let i = 0; i < urls.length; i++) {
      imgHtml += `<div><img style="width: 200px" src="/${urls[i]}" /> <div>/${urls[i]}</div></div>`;
    }

    $(".file_upload-msg").html(
      `<div>${new Date()}<div>
        ${imgHtml}
        </div>
    </div>`
    );
});
```

[官方文件上传 demo](https://github.com/gin-gonic/examples/tree/master/upload-file)

## 文件分片上传原理

客户端会根据文件大小和用户要分片的大小来计算文件分片个数。客户端会一片一片的去请求接口把文件的所有片段上传带服务器端。

服务端接受客户端上传的文件片段进行缓存或创建文件并读入该片段，直至最后一片上传成功。

> [http://localhost:8080/upload_chunks](http://localhost:8080/upload_chunks)

## 服务器端

服务器端使用的是 go 语言的 gin 框架。

```go
type ChunkFile struct {
	Name   string         `json:"name" form:"name"`
	Chunk  int            `json:"chunk" form:"chunk"`
	Chunks int            `json:"chunks" form:"chunks"`
}

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
// 文件分片上传handler
func fileChunkUpload(c *gin.Context) {

	var chunkFile ChunkFile
	r := c.Request

	c.Bind(&chunkFile)

	var Buf = make([]byte, 0)
	// in your case file would be fileupload
	file, _, _ := r.FormFile("file")

	log.Println("this is ", chunkFile.File)
	Buf, _ = ioutil.ReadAll(file)

	filePath := "upload/"+ chunkFile.Name

	fd, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.Write(Buf)
	fd.Close()

	if chunkFile.Chunk + 1 == chunkFile.Chunks {
		c.JSON(http.StatusOK, gin.H{
			"state": "SUCCESS",
			"url": "/"+filePath,
		})
	} else {
		contentType := strings.Split(c.GetHeader("Content-Type"), "boundary=")
		c.String(http.StatusOK, contentType[1])
	}
}
```

[服务端接口完整代码](https://github.com/freeshineit/gin_rotuer_web/blob/master/api)

## 客户端（web）

客户端使用的是[plupload](https://www.plupload.com/)文件上传插件，好处是它提供了分片上传，创建对象时配置`chunk_size`属性就可以实现了（插件底层会根据文件大小和`chunk_size`来计算分片的个数）。

```js
var uploader = new plupload.Uploader({
  runtimes: "html5,flash,silverlight,html4",
  browse_button: "pickfiles", // you can pass an id...
  container: document.getElementById("container"), // ... or DOM Element itself
  url: "/api/file_chunk_upload",
  flash_swf_url: "/static/js/Moxie.swf",
  silverlight_xap_url: "/static/js/Moxie.xap",
  chunk_size: "100kb",
  filters: {
    max_file_size: "10mb",
    mime_types: [
      { title: "Image files", extensions: "jpg,gif,png,jpeg" },
      { title: "Zip files", extensions: "zip" },
    ],
  },

  init: {
    PostInit: function () {
      document.getElementById("filelist").innerHTML = "";

      document.getElementById("uploadfiles").onclick = function () {
        uploader.start();
        return false;
      };
    },

    FilesAdded: function (up, files) {
      plupload.each(files, function (file) {
        document.getElementById("filelist").innerHTML +=
          '<div id="' +
          file.id +
          '">' +
          file.name +
          " (" +
          plupload.formatSize(file.size) +
          ") <b></b></div>";
      });
    },

    UploadProgress: function (up, file) {
      document.getElementById(file.id).getElementsByTagName("b")[0].innerHTML =
        "<span>" + file.percent + "%</span>";
    },

    Error: function (up, err) {
      document
        .getElementById("console")
        .appendChild(
          document.createTextNode("\nError #" + err.code + ": " + err.message)
        );
    },
  },
});

uploader.bind("ChunkUploaded", function (up, file, info) {
  // do some chunk related stuff
  console.log(info);
});

uploader.init();
```

[客户端文件上传完整代码](https://github.com/freeshineit/gin_rotuer_web/blob/master/templates/upload_chunks.html)

[demo](https://github.com/freeshineit/gin_rotuer_web)
