# gin_router_web

![build](https://github.com/freeshineit/gin_rotuer_web/workflows/build/badge.svg)

[gin](https://github.com/gin-gonic/gin)是简单快速的`golang`框架，这个项目主要是介绍`gin`的路由配置及使用，包括各种HTTP请求方法（GET、POST）和数据格式处理（JSON、XML、Form、URL Encoded）。

**主要特性：**
- 展示 Gin 框架的多种路由配置方式
- 支持多种数据格式的请求处理（JSON、XML、Form、URL Encoded）
- 文件上传功能（单文件及分片上传）
- 前端与后端交互示例

**技术栈：**
- Go >= 1.18
- Gin Web Framework v1.9.1
- 支持 Docker 部署



## 使用

```bash
# development
go run main.go

# run development with live reload
# https://github.com/cosmtrek/air Live reload for Go apps
air

# build
go build

# run production
# export GIN_MODE=release
./gin-router-web

# make build -> ./bin/app
make build

# server 8080
http://localhost:8080/

# file chunk upload
http://localhost:8080/upload_chunks

# docker deploy
make serve

# dependency management
go mod tidy
```

## 项目理念与设计

本项目是一个 **Gin 框架学习和示例项目**，旨在展示：

### 核心特性
1. **路由配置** - 展示 Gin 框架的路由分组、中间件、静态资源配置等
2. **多种数据格式处理** - JSON、XML、Form、URL Encoded 等多种数据格式的处理方式
3. **文件处理** - 单文件上传、批量上传、分片上传完整解决方案
4. **Web 应用** - 包含前后端交互的完整示例

### 架构思想
- **分层设计** - API层、路由层、模型层清晰分离
- **可扩展性** - 新增 API 只需在 `api/` 目录添加处理函数，在 `route.go` 注册路由
- **代码复用** - 响应格式统一通过 `helper` 模块处理
- **前后端分离** - 前端页面独立，便于前后端开发解耦

## 项目结构

```
gin_router_web/
├── api/                   # API 处理层
│   ├── handle_func.go    # 各类型数据处理函数
│   ├── upload.go         # 文件上传相关
│   └── web.go            # 网页路由处理
├── router/
│   └── route.go          # 路由配置
├── models/               # 数据模型
│   ├── chunk_file.go     # 分片文件模型
│   └── user.go           # 用户模型
├── helper/
│   └── respose.go        # 响应格式化工具
├── templates/            # HTML 模板
│   ├── index.html        # 首页
│   └── upload_chunks.html # 分片上传页面
├── public/               # 静态资源
│   └── static/
│       ├── css/
│       │   └── common.css
│       └── js/
│           ├── index.js
│           └── plupload.full.min.js
├── main.go              # 应用入口
├── go.mod               # Go 模块定义
├── Dockerfile           # Docker 配置
├── docker-compose.yml   # Docker Compose 配置
├── Makefile            # 构建脚本
└── README.md           # 项目文档
```

### 关键模块说明

| 模块 | 描述 |
|-----|------|
| **api** | 包含所有的 HTTP 请求处理函数，支持多种数据格式 |
| **router** | 路由配置层，定义所有的 API 端点和Web路由 |
| **models** | 数据模型定义，如用户信息、分片文件等 |
| **helper** | 工具函数，如统一的响应格式化 |
| **templates** | HTML 模板文件，提供前端界面 |
| **public** | 静态资源（CSS、JavaScript等） |

## API 接口列表

```md
[GET]    /                       首页
[GET]    /upload_chunks          分片上传页面
[GET]    /api/query              查询参数示例

[POST]   /api/form_post          表单数据提交
[POST]   /api/json_post          JSON 数据提交
[POST]   /api/urlencoded_post    URL 编码数据提交
[POST]   /api/json_and_form_post JSON 或 Form 数据提交（自动识别）
[POST]   /api/xml_post           XML 数据提交
[POST]   /api/file_upload        文件上传
[POST]   /api/file_chunk_upload  分片文件上传
```

## 快速开始

### 环境要求
- Go >= 1.18
- (可选) Docker & Docker Compose 用于容器化部署

### 本地运行

```bash
# 1. 克隆项目
git clone https://github.com/freeshineit/gin_router_web.git
cd gin_router_web

# 2. 安装依赖
go mod tidy

# 3. 开发模式运行（需要安装 air）
go install github.com/cosmtrek/air@latest
air

# 或直接运行
go run main.go

# 4. 浏览器访问
http://localhost:8080/
```

### Docker 部署

```bash
# 构建并运行 Docker 容器
make deploy

# 或手动构建
docker build -t gin-router-web .
docker run -p 8080:8080 gin-router-web
```

## 静态资源配置

```go
func setStaticFS(r *gin.Engine) {
	// 加载 HTML 模板文件
	r.LoadHTMLGlob("./templates/*.html")

	// 配置 favicon
	r.StaticFile("favicon.ico", "./public/favicon.ico")
	
	// 配置静态文件目录（CSS、JS 等）
	r.StaticFS("/static", http.Dir("public/static"))
	
	// 配置上传文件目录
	r.StaticFS("/upload", http.Dir("upload"))
}
```

**关键方法说明：**

- `LoadHTMLGlob(pattern string)` - 加载全局模式的 HTML 文件标识，并与 HTML 渲染器关联
- `StaticFile(relativePath, filepath string)` - 配置单个静态文件
- `StaticFS(relativePath string, fs http.FileSystem)` - 配置静态资源目录

## 路由配置详解

### Web 路由

```go
func setWebRoute(r *gin.Engine) {
	// 首页路由
	r.GET("/", api.WebIndex)
	
	// 分片上传页面
	r.GET("/upload_chunks", api.WebUploadChunks)
}
```

### API 路由分组

```go
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// 设置静态资源
	setStaticFS(r)

	// 设置Web路由
	setWebRoute(r)

	// API 路由分组
	apiGroup := r.Group("/api")
	{
		// 表单提交
		apiGroup.POST("/form_post", api.FormPost)

		// JSON 提交
		apiGroup.POST("/json_post", api.JSONPost)

		// URL 编码提交
		apiGroup.POST("/urlencoded_post", api.UrlencodedPost)

		// JSON 和 Form 混合提交（自动识别）
		apiGroup.POST("/json_and_form_post", api.JSONAndFormPost)

		// XML 提交
		apiGroup.POST("/xml_post", api.XMLPost)

		// 文件上传
		apiGroup.POST("/file_upload", api.FileUpload)

		// 文件分片上传
		apiGroup.POST("/file_chunk_upload", api.FileChunkUpload)

		// 查询参数示例
		apiGroup.GET("/query", func(c *gin.Context) {
			name := c.Query("name")
			message := c.Query("message")
			nick := c.DefaultQuery("nick", "anonymous")

			c.JSON(http.StatusOK, helper.BuildResponse(gin.H{
				"name":    name,
				"message": message,
				"nick":    nick,
			}))
		})
	}

	return r
}
```

**关键概念：**
- **路由分组** (`r.Group()`) - 对多个相关路由进行分组，便于批量配置中间件或路径前缀
- **Query 参数** - 通过 URL 查询字符串传递，如 `?name=value&message=test`
- **DefaultQuery** - 获取查询参数，提供默认值

## HTTP 请求数据格式

在 HTTP 通信中，请求体的数据格式由 `Content-Type` 头部指定。本项目展示如何处理常见的数据格式：

| Content-Type | 用途 | 例子 |
|-------------|------|-----|
| `text/plain` | 纯文本 | 简单文本消息 |
| `text/html` | HTML 文档 | 网页内容 |
| `application/json` | JSON 格式数据 | REST API 的标准格式 |
| `application/x-www-form-urlencoded` | 表单数据 | HTML 表单默认格式 |
| `application/xml` | XML 格式数据 | 配置文件、SOAP 通信 |
| `multipart/form-data` | 多部分数据 | 文件上传、混合数据 |

参考：[MIME 类型](https://zh.wikipedia.org/wiki/MIME)

### 1. Form 表单提交 (`application/x-www-form-urlencoded`)

**数据模型**

```go
// User 用户信息结构体，支持 JSON、Form 和 XML 三种格式
type User struct {
	Name    string `json:"name" form:"name" xml:"name"`
	Message string `json:"message" form:"message" xml:"message"`
	Nick    string `json:"nick" form:"nick" xml:"nick"`
}
```

**后端实现** (api/handle_func.go)

```go
// FormPost 处理表单数据提交
func FormPost(c *gin.Context) {
	// 方式1：逐个读取字段
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "default nick")
	name := c.DefaultPostForm("name", "default name")
	
	user := User{
		Name:    name,
		Nick:    nick,
		Message: message,
	}

	// 方式2：自动绑定（推荐）
	// user := &User{}
	// c.ShouldBind(user)

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

**前端实现** (templates/index.html)

```html
<form method="post" action="/api/form_post" id="form">
  <div class="form-item">
    <label for="name">名字</label>
    <input type="text" id="name" name="name" />
  </div>
  <div class="form-item">
    <label for="message">消息</label>
    <input type="text" id="message" name="message" />
  </div>
  <div class="form-item">
    <label for="nick">昵称</label>
    <input type="text" id="nick" name="nick" />
  </div>
  <button type="submit">提交</button>
</form>
```

**关键知识点：**
- `c.PostForm(key)` - 读取表单字段
- `c.DefaultPostForm(key, defaultValue)` - 读取表单字段，提供默认值
- `c.ShouldBind()` - 自动将表单数据绑定到结构体（推荐）

### 2. JSON 数据提交 (`application/json`)

**后端实现**

```go
// JSONPost 处理 JSON 格式数据
func JSONPost(c *gin.Context) {
	var user User
	
	// 绑定 JSON 数据到结构体
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(
			http.StatusBadRequest, 
			"invalid parameter",
		))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

**前端实现** (JavaScript/Axios)

```js
const data = {
  name: "shine",
  message: "hello gin",
  nick: "shineshao"
};

axios({
  method: "post",
  url: "/api/json_post",
  headers: {
    "Content-Type": "application/json",
  },
  data, // 自动序列化为 JSON
}).then((res) => {
  console.log(res.data);
  $(".json-msg").text(`success at ${new Date()}`);
});
```

**关键知识点：**
- `c.BindJSON()` - 解析 JSON 请求体并绑定到结构体
- Axios 会自动设置 `Content-Type: application/json`

### 3. URL 编码数据提交 (`application/x-www-form-urlencoded`)

**后端实现**

```go
// UrlencodedPost 处理 URL 编码格式数据
func UrlencodedPost(c *gin.Context) {
	// 从查询参数读取 limit
	limit := c.Query("limit")
	
	// 从 POST 表单数据读取
	name := c.PostForm("name")
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "default")
	
	user := User{
		Name:    name,
		Nick:    nick,
		Message: message,
	}

	log.Printf("request query limit: %s\n", limit)
	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

**前端实现**

```js
const data = {
  name: "shine",
  message: "hello gin",
  nick: "shineshao"
};

// 注意：查询参数在 URL 中，表单数据在请求体中
axios({
  method: "post",
  url: "/api/urlencoded_post?limit=100",
  headers: {
    "Content-Type": "application/x-www-form-urlencoded",
  },
  data: $.param(data), // 使用 $.param() 进行 URL 编码
}).then((res) => {
  console.log(res.data);
  $(".urlencoded-msg").text(`success at ${new Date()}`);
});
```

**关键知识点：**
- `c.Query()` - 读取 URL 查询参数
- `c.PostForm()` - 读取 POST 表单数据
- 使用 `$.param()` 进行 URL 编码

### 4. JSON 或 Form 混合提交 (自动识别)

**后端实现**

```go
// JSONAndFormPost 自动识别 JSON 或 Form 格式
// 可以同时处理 application/json 和 application/x-www-form-urlencoded
func JSONAndFormPost(c *gin.Context) {
	var user User

	// ShouldBind() 根据 Content-Type 自动选择合适的绑定方式
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(
			http.StatusBadRequest, 
			"invalid parameter",
		))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

**前端实现 - JSON 方式**

```js
const data = {
  name: "shine",
  message: "hello gin",
  nick: "shineshao"
};

axios({
  method: "post",
  url: "/api/json_and_form_post",
  headers: {
    "Content-Type": "application/json",
  },
  data,
}).then((res) => {
  console.log(res.data);
  $(".jsonandform-msg").text(`success with JSON at ${new Date()}`);
});
```

**前端实现 - Form 方式**

```js
axios({
  method: "post",
  url: "/api/json_and_form_post",
  headers: {
    "Content-Type": "application/x-www-form-urlencoded",
  },
  data: $.param(data),
}).then((res) => {
  console.log(res.data);
  $(".jsonandform-msg").text(`success with Form at ${new Date()}`);
});
```

**关键知识点：**
- `c.ShouldBind()` - 自动识别 Content-Type，使用相应的绑定方式
- 同一个接口可以处理多种数据格式

### 5. XML 数据提交 (`application/xml`)

**后端实现**

```go
// XMLPost 处理 XML 格式数据
func XMLPost(c *gin.Context) {
	var user User

	// 绑定 XML 数据到结构体
	if err := c.BindXML(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(
			http.StatusBadRequest, 
			"invalid parameter",
		))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(user))
}
```

**前端实现**

```js
const data = {
  name: "shine",
  message: "hello gin",
  nick: "shineshao"
};

// 手动构建 XML 字符串
const xmlData = `<xml>
  <name>${data.name}</name>
  <message>${data.message}</message>
  <nick>${data.nick}</nick>
</xml>`;

axios({
  method: "post",
  url: "/api/xml_post",
  headers: {
    "Content-Type": "application/xml",
  },
  data: xmlData,
}).then((res) => {
  console.log(res.data);
  $(".xml-msg").text(`success at ${new Date()}`);
});
```

**关键知识点：**
- `c.BindXML()` - 解析 XML 请求体并绑定到结构体
- 需要在结构体标签中指定 `xml:"fieldname"` 标签
- XML 通常用于旧系统集成、配置文件等场景

### 6. 文件上传 (`multipart/form-data`)

**数据模型**

```go
// ChunkFile 分片文件信息
type ChunkFile struct {
	Name   string `json:"name" form:"name"`       // 文件名
	Chunk  int    `json:"chunk" form:"chunk"`     // 当前分片号
	Chunks int    `json:"chunks" form:"chunks"`   // 总分片数
}
```

**后端实现** (api/upload.go)

```go
// FileUpload 处理文件上传
func FileUpload(c *gin.Context) {
	filesUrl := make([]string, 0)
	
	// 获取多部分表单数据
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("MultipartForm error: %s", err)
		return
	}

	// 获取所有 "file" 字段的文件
	files := form.File["file"]

	// 创建 upload 目录（如果不存在）
	if _, err := os.Stat("upload"); err != nil {
		os.Mkdir("upload", os.ModePerm)
	}

	// 保存所有上传的文件
	for _, file := range files {
		log.Println("Uploading file:", file.Filename)

		// 保存文件到指定位置
		if err := c.SaveUploadedFile(file, "upload/"+file.Filename); err != nil {
			log.Println("SaveUploadedFile error: %s", err)
			return
		}
		filesUrl = append(filesUrl, "upload/"+file.Filename)
	}

	c.JSON(http.StatusOK, helper.BuildResponse(gin.H{
		"urls": filesUrl,
	}))
}
```

**前端实现 - HTML**

```html
<div>
  <form id="multipleForm">
    <!-- multiple 属性支持多文件选择 -->
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

**前端实现 - JavaScript**

```js
// 单个文件上传
// const fd = new FormData()
// const file = document.getElementById('file')
// fd.append('file', file.files[0])

// 多个文件上传
axios({
  method: "post",
  url: "/api/file_upload",
  headers: {
    "Content-Type": "multipart/form-data",
  },
  data: new FormData($("#multipleForm")[0]),
}).then((res) => {
  console.log(res.data);
  const urls = res.data.data.urls || [];

  let imgHtml = "";
  for (let i = 0; i < urls.length; i++) {
    imgHtml += `<div>
      <img style="width: 200px" src="/${urls[i]}" /> 
      <div>${urls[i]}</div>
    </div>`;
  }

  $(".file_upload-msg").html(`<div>${new Date()}</div>${imgHtml}`);
});
```

**关键知识点：**
- `c.MultipartForm()` - 获取多部分表单数据
- `c.SaveUploadedFile()` - 保存上传的文件
- `FormData` - JavaScript 用于构建表单数据（包括文件）
- 使用 `multiple` 属性支持多文件选择

**参考资源：** [Gin 官方文件上传示例](https://github.com/gin-gonic/examples/tree/master/upload-file)

## 文件分片上传

### 工作原理

文件分片上传是一种大文件上传优化技术，通过将大文件拆分成多个小片段，分别上传到服务器：

1. **客户端** - 根据文件大小和设定的分片大小，计算需要的分片数量
2. **分片上传** - 逐片发送分片数据到服务器
3. **服务端** - 接收并缓存各片段，当所有片段都收到后，完成文件上传

**优势：**
- 支持断点续传
- 网络不稳定时可单独重试失败的分片
- 提高大文件上传的成功率和用户体验

### 服务器实现

**数据模型** (models/chunk_file.go)

```go
type ChunkFile struct {
	Name   string `json:"name" form:"name"`       // 文件名
	Chunk  int    `json:"chunk" form:"chunk"`     // 当前分片编号（0开始）
	Chunks int    `json:"chunks" form:"chunks"`   // 总分片数
}
```

**处理函数** (api/upload.go)

```go
// FileChunkUpload 处理文件分片上传
func FileChunkUpload(c *gin.Context) {
	var chunkFile ChunkFile
	r := c.Request

	// 绑定分片信息
	c.Bind(&chunkFile)

	// 读取上传的文件数据
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("FormFile error:", err)
		return
	}

	// 将文件内容读入缓冲区
	buf, _ := ioutil.ReadAll(file)

	// 文件保存路径
	filePath := "upload/" + chunkFile.Name

	// 打开文件，以追加模式写入（如果不存在则创建）
	fd, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.Write(buf)
	fd.Close()

	// 判断是否是最后一片
	if chunkFile.Chunk+1 == chunkFile.Chunks {
		// 所有分片已上传完成
		c.JSON(http.StatusOK, gin.H{
			"state": "SUCCESS",
			"url":   "/" + filePath,
		})
	} else {
		// 还有更多分片待上传
		c.String(http.StatusOK, "UPLOADING")
	}
}
```

### 客户端实现

项目使用 [plupload](https://www.plupload.com/) 插件实现分片上传。该插件自动处理文件分片、并发上传等细节。

**初始化配置** (templates/upload_chunks.html)

```js
var uploader = new plupload.Uploader({
	runtimes: "html5,flash,silverlight,html4",
	browse_button: "pickfiles",
	container: document.getElementById("container"),
	url: "/api/file_chunk_upload",
	flash_swf_url: "/static/js/Moxie.swf",
	silverlight_xap_url: "/static/js/Moxie.xap",
	
	// 分片大小：100KB
	chunk_size: "100kb",
	
	// 文件过滤
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
			// 文件添加事件
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
			// 上传进度事件
			document.getElementById(file.id).getElementsByTagName("b")[0].innerHTML =
				"<span>" + file.percent + "%</span>";
		},

		Error: function (up, err) {
			// 错误事件
			document
				.getElementById("console")
				.appendChild(
					document.createTextNode(
						"\nError #" + err.code + ": " + err.message
					)
				);
		},

		ChunkUploaded: function (up, file, info) {
			// 单个分片上传完成事件
			console.log("Chunk uploaded:", file.name, info);
		},
	},
});

uploader.init();
```

**关键配置参数说明：**

| 参数 | 说明 |
|-----|------|
| `url` | 分片上传的服务器接口地址 |
| `chunk_size` | 分片大小（KB 或 MB） |
| `max_file_size` | 允许的最大文件大小 |
| `runtimes` | 支持的运行环境 |
| `FilesAdded` | 文件选择后的回调 |
| `UploadProgress` | 上传进度回调 |
| `ChunkUploaded` | 分片上传完成回调 |

**演示页面：** [http://localhost:8080/upload_chunks](http://localhost:8080/upload_chunks)

**参考资源：**
- [服务端接口完整代码](https://github.com/freeshineit/gin_rotuer_web/blob/master/api)
- [客户端文件上传完整代码](https://github.com/freeshineit/gin_rotuer_web/blob/master/templates/upload_chunks.html)
- [完整项目演示](https://github.com/freeshineit/gin_rotuer_web)

## 最佳实践

### 1. 数据绑定

```go
// ✅ 推荐：使用 ShouldBind，错误处理更灵活
var user User
if err := c.ShouldBind(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

// ❌ 不推荐：使用 Bind，当绑定失败时会自动响应 400
c.Bind(&user)
```

### 2. 响应格式统一

```go
// 使用 helper 函数统一响应格式
c.JSON(http.StatusOK, helper.BuildResponse(data))
c.JSON(http.StatusBadRequest, helper.BuildErrorResponse(400, "error message"))
```

### 3. 文件上传安全性

```go
// ✅ 验证文件类型
allowedTypes := map[string]bool{
    "image/jpeg": true,
    "image/png":  true,
    "image/gif":  true,
}

fileHeader, _ := c.FormFile("file")
if !allowedTypes[fileHeader.Header.Get("Content-Type")] {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type"})
    return
}

// ✅ 限制文件大小
if fileHeader.Size > 10*1024*1024 { // 10MB
    c.JSON(http.StatusBadRequest, gin.H{"error": "file too large"})
    return
}
```

### 4. 错误处理

```go
// 使用 defer 统一处理错误恢复
defer func() {
    if r := recover(); r != nil {
        log.Printf("Panic recovered: %v", r)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
    }
}()
```

### 5. 日志记录

```go
// 记录关键操作
log.Printf("User: %s, Action: upload, Size: %d bytes, Time: %s",
    username,
    fileHeader.Size,
    time.Now().Format("2006-01-02 15:04:05"),
)
```

## 常见问题

### Q1：上传文件时出现 "invalid request" 错误？
**A：** 检查 `Content-Type` 头是否正确设置为 `multipart/form-data`。使用 Axios 的 FormData 时应该不设置 Content-Type，让浏览器自动处理。

```js
// ✅ 正确
const formData = new FormData();
formData.append('file', file);
axios.post('/api/file_upload', formData); // 不设置 headers

// ❌ 错误
axios.post('/api/file_upload', formData, {
    headers: { 'Content-Type': 'application/json' }
});
```

### Q2：如何处理 CORS 跨域问题？
**A：** 使用 Gin 的 CORS 中间件：

```bash
go get github.com/gin-contrib/cors
```

```go
import "github.com/gin-contrib/cors"

r := gin.Default()
r.Use(cors.Default())
```

### Q3：大文件上传超时如何解决？
**A：** 增加服务器的请求超时时间：

```go
srv := &http.Server{
    Addr:         ":8080",
    Handler:      r,
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
    IdleTimeout:  60 * time.Second,
}
```

### Q4：如何保证文件上传的安全性？
**A：** 采取以下措施：
- 验证文件类型（通过 MIME 类型和文件扩展名）
- 限制文件大小
- 使用白名单过滤文件名中的特殊字符
- 存储文件到不可直接访问的目录
- 使用 HTTPS 传输

```go
// 文件名清理示例
fileName := filepath.Base(fileHeader.Filename)
fileName = strings.Map(func(r rune) rune {
    if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' || r == '_' || r == '-' {
        return r
    }
    return -1
}, fileName)
```

## 学习资源

- **Gin 官方文档**：https://gin-gonic.com/
- **Go 标准库文档**：https://golang.org/pkg/
- **HTTP 协议规范**：https://tools.ietf.org/html/rfc7231
- **Plupload 文档**：https://www.plupload.com/

## 扩展建议

1. **数据库集成** - 添加 GORM 或其他 ORM，持久化用户数据
2. **身份验证** - 实现 JWT 或 Session 认证机制
3. **日志系统** - 集成 Logrus 或 Zap 提供结构化日志
4. **限流控制** - 实现请求限流，防止滥用
5. **单元测试** - 为 API 编写完整的单元测试
6. **错误监控** - 集成 Sentry 或其他错误监控服务
7. **性能优化** - 添加缓存、CDN 支持
8. **文档自动生成** - 使用 Swag 自动生成 OpenAPI 文档

## 许可证

本项目采用 MIT 许可证。详见 LICENSE 文件。
