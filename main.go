package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wuttinanhi/go-web-file-upload/config"
	"github.com/wuttinanhi/go-web-file-upload/route"
	"github.com/wuttinanhi/go-web-file-upload/util"
)

func main() {
	// create upload dir
	util.CreateDir(config.GetConfig().UPLOAD_DESTINATION)

	// create gin router
	router := gin.Default()

	// recovery middleware
	router.Use(gin.Recovery())

	// load template
	router.LoadHTMLGlob("templates/*")

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// expose upload directory
	router.Static("/upload", config.GetConfig().UPLOAD_DESTINATION)

	// index route
	router.GET("/", route.IndexRoute)

	// file upload route
	router.POST("/upload/single", route.UploadRoute)

	// gallery route
	router.GET("/gallery", route.GalleryRoute)

	// run server
	router.Run(":8080")
}
