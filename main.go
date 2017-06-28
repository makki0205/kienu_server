package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/HALDevelopersTeam/crow_server/middleware"
	"github.com/HALDevelopersTeam/crow_server/controller"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "yes")
	})

	r.Use()
	r.GET("/file/:path/:name", func(c *gin.Context){
		path := c.Param("path")
		name := c.Param("name")
		c.File("./storage/file/" + path + "/" + name)
	})
	api := r.Group("/api")
	api.Use(cors.Middleware(middleware.CorsConfig))

	fctr := controller.NewFileCtr()
	api.POST("/upload", fctr.UploadFile)
	api.GET("/@:uuid", fctr.GetFileDescription)

	r.Run(":3000")
}