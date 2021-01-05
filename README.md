# hotel_cms
软工课设，使用go语言粗略实现一个酒店管理系统。

[风格指南](https://www.oschina.net/news/110525/uber-go-style-guide)

## 需求分析

### 人事工资管理

在职人员的管理，身份管理，包括总经理，前台员工，保安，清洁工。酒店副总、各部门经理、各部门总监

[等级](https://wenku.baidu.com/view/f3d03db6ad51f01dc381f123.html)

### 权限表





### 考勤系统



月表



### 房卡管理系统

房卡的注册、房卡开卡、消卡、挂失等。

### 客房管理

​	清理完毕空闲状态可预订、 入住、 入住清理、 免打扰、 退房中待验收、验收完毕待清理、清理完毕空闲状态可预订  预定人/入住人  人数



### 房间类型管理

类型、基本价格、数量、入住人数



#### 房间订单

房间号，客户姓名，手机号，开始时间，结束时间



### 黑名单管理

出现了迟退、房间损坏严重、超时三次



### 入住记录

身份证加密



### 活动

常价、滞销折扣、节假日涨价



### 日用品库存管理

被子、枕头、床单、桌子、椅子、牙刷、梳子、吹风机、卷纸、拖鞋、衣架、毛巾、浴巾、洗发液、沐浴露、矿泉水、电视......

入库、出库、记明细帐、总帐、盘存、分仓管理、报表。

### 旅行社、合作方（经理）

几个阶段：小型，中型，大型，折扣大小

### 开销报表


## gin框架

### 学习

```go
package main
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```

1. 使用gin.Default()生成一个实例，这个实例即 WSGI 应用程序。
2. r.Get("/", ...)声明了一个路由，告诉 Gin 什么样的URL 能触发传入的函数，
这个函数返回我们想要显示在用户浏览器中的信息。
3. r.Run()函数让应用运行在本地服务器上，默认监听端口是 _8080_，
   可以传入参数设置端口，例如r.Run(":9999")即运行在 _9999_端口。


### gin路由(Route)

1. 路由方法有 GET, POST, PUT, PATCH, DELETE 和 OPTIONS，还有Any，
可匹配以上任意类型的请求。

```go
// GET 无参数
r.GET("/", func(c *gin.Context) {
	c.String(http.StatusOK, "Who are you?")
})


// GET 带参数 
// 匹配users?name=xxx&role=xxx，role可选
r.GET("/users", func(c *gin.Context) {
name := c.Query("name")
role := c.DefaultQuery("role", "teacher")
c.String(http.StatusOK, "%s is a %s", name, role)
})

// POST
// 'username=geektutu&password=1234'
// {"password":"1234","username":"geektutu"}
r.POST("/form", func(c *gin.Context) {
    username := c.PostForm("username")
    password := c.DefaultPostForm("password", "000000") // 可设置默认值
    
    c.JSON(http.StatusOK, gin.H{
    "username": username,
    "password": password,
    })
})

// GET 和 POST 混合
// curl "http://localhost:9999/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
// {"id":"9876","page":"7","password":"1234","username":"geektutu"}
r.POST("/posts", func(c *gin.Context) {
    id := c.Query("id")
    page := c.DefaultQuery("page", "0")
    username := c.PostForm("username")
    password := c.DefaultPostForm("username", "000000") // 可设置默认值
    
    c.JSON(http.StatusOK, gin.H{
    "id":       id,
    "page":     page,
    "username": username,
    "password": password,
    })
})

// post map传参
// curl -g "http://localhost:9999/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
// {"ids":{"Jack":"001","Tom":"002"},"names":{"a":"Sam","b":"David"}}
r.POST("/post", func(c *gin.Context) {
    ids := c.QueryMap("ids")
    names := c.PostFormMap("names")
    
    c.JSON(http.StatusOK, gin.H{
    "ids":   ids,
    "names": names,
    })
})

// POST JSON 
engine.PUT("/user", func(context *gin.Context) {
   var m map[string]string
   if err := context.BindJSON(&m); err != nil {
       context.String(http.StatusInternalServerError, "error data!")
       return
   }
   context.String(http.StatusOK, "hello " + m["name"])
})

// 重定向的两种方式

r.GET("/redirect", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "/index")
})

r.GET("/goIndex", func(c *gin.Context) {
    c.Request.URL.Path = "/"
    r.HandleContext(c)
})

// route 分组
// group routes 分组路由
defaultHandler := func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "path": c.FullPath(),
    })
}
// group: v1
v1 := r.Group("/v1"){
    v1.GET("/posts", defaultHandler)
    v1.GET("/series", defaultHandler)
}
// group: v2
v2 := r.Group("/v2"){
    v2.GET("/posts", defaultHandler)
    v2.GET("/series", defaultHandler)
}


//  upload 单个文件
r.POST("/upload1", func(c *gin.Context) {
    file, _ := c.FormFile("file")
    // c.SaveUploadedFile(file, dst)
    c.String(http.StatusOK, "%s uploaded!", file.Filename)
})
// 多个文件
r.POST("/upload2", func(c *gin.Context) {
    // Multipart form
    form, _ := c.MultipartForm()
    files := form.File["upload[]"]
    
    for _, file := range files {
        log.Println(file.Filename)
        // c.SaveUploadedFile(file, dst)
    }
    c.String(http.StatusOK, "%d files uploaded!", len(files))
})


```