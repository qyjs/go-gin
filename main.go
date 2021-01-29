package main

import (
	"fmt"
	db "gin/databases"
	"time"
)

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	// 给表单限制上传大小 (默认 32 MiB)
// 	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
// 	router.POST("/upload", func(c *gin.Context) {
// 		// 单文件
// 		file, _ := c.FormFile("file")
// 		log.Println(file.Filename)

// 		// 上传文件到指定的路径
// 		// c.SaveUploadedFile(file, dst)

// 		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
// 	})
// 	router.Run(":8080")
// }

// Login ...
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	defer db.SQLDB.Close()

	router := initRouter()
	router.Run(":8000")

	// router := gin.Default()

	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }
	// router.LoadHTMLGlob("templates/*")
	// //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	// router.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Main website",
	// 	})
	// })

	// router.Delims("{[{", "}]}")
	// router.SetFuncMap(template.FuncMap{
	// 	"formatAsDate": formatAsDate,
	// })
	// router.LoadHTMLFiles("./testdata/template/raw.tmpl")

	// router.GET("/raw", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
	// 		"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
	// 	})
	// })

	// router.GET("/cookie", func(c *gin.Context) {
	// 	cookie, err := c.Cookie("gin_cookie")
	// 	if err != nil {
	// 		cookie = "NotSet"
	// 		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	// 	}
	// 	fmt.Printf("Cookie value: %s \n", cookie)
	// })

	// // Serves unicode entities
	// router.GET("/json", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"html": "<b>Hello, world!</b>",
	// 	})
	// })

	// // Serves literal characters
	// router.GET("/purejson", func(c *gin.Context) {
	// 	c.PureJSON(200, gin.H{
	// 		"html": "<b>Hello, world!</b>",
	// 	})
	// })

	// router.GET("/someDataFromReader", func(c *gin.Context) {
	// 	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	// 	if err != nil || response.StatusCode != http.StatusOK {
	// 		c.Status(http.StatusServiceUnavailable)
	// 		return
	// 	}

	// 	reader := response.Body
	// 	contentLength := response.ContentLength
	// 	contentType := response.Header.Get("Content-Type")

	// 	extraHeaders := map[string]string{
	// 		"Content-Disposition": `attachment; filename="gopher.png"`,
	// 	}

	// 	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	// })

	// // Example for binding JSON ({"user": "manu", "password": "123"})
	// router.POST("/loginJSON", func(c *gin.Context) {
	// 	var json Login
	// 	if err := c.ShouldBindJSON(&json); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	if json.User != "manu" || json.Password != "123" {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	// })

	// // Example for binding XML (
	// //  <?xml version="1.0" encoding="UTF-8"?>
	// //  <root>
	// //      <user>user</user>
	// //      <password>123</password>
	// //  </root>)
	// router.POST("/loginXML", func(c *gin.Context) {
	// 	var xml Login
	// 	if err := c.ShouldBindXML(&xml); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	if xml.User != "manu" || xml.Password != "123" {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	// })

	// // Example for binding a HTML form (user=manu&password=123)
	// router.POST("/loginForm", func(c *gin.Context) {
	// 	var form Login
	// 	if err := c.ShouldBind(&form); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	if form.User != "manu" || form.Password != "123" {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	// })
	// router.Run(":8080")

}
