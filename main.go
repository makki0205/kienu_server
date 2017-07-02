package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/makki0205/tmp.fun/middleware"
	"github.com/makki0205/tmp.fun/controller"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")
	r.LoadHTMLGlob("view/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	fctr := controller.NewFileCtr()
	r.GET("/@:uuid", fctr.GetFile)
	api := r.Group("/api")
	api.Use(cors.Middleware(middleware.CorsConfig))
	api.OPTIONS("/:hoge", func(c *gin.Context) {

	})
	api.POST("/upload", fctr.UploadFile)
	api.GET("/@:uuid", fctr.GetFileDescription)

	r.Run(":3000")
}