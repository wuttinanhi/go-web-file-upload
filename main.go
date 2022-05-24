package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

// upload destination
var UPLOAD_DESTINATION = path.Join("public", "upload")

var ALLOWED_CONTENT_TYPE = []string{"image/jpeg", "image/png", "image/gif", "image/bmp"}

func main() {
	// create upload dir
	CreateDir(UPLOAD_DESTINATION)

	// create gin router
	router := gin.Default()

	// recovery middleware
	router.Use(gin.Recovery())

	// load template
	router.LoadHTMLGlob("templates/*")

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// expose upload directory
	router.Static("/upload", UPLOAD_DESTINATION)

	// index route
	router.GET("/", IndexRoute)
	// file upload route
	router.POST("/upload/single", UploadRoute)
	// gallery route
	router.GET("/gallery", GalleryRoute)

	// run server
	router.Run(":8080")
}

func CreateDir(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModePerm)
	}
}

func IsAllowedContentType(contentType string) bool {
	for _, allowedType := range ALLOWED_CONTENT_TYPE {
		if allowedType == contentType {
			return true
		}
	}
	return false
}

func IndexRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func UploadRoute(c *gin.Context) {
	// Single file
	file, _ := c.FormFile("file")

	// read file as byte
	data, _ := file.Open()
	defer data.Close()
	fileBytes, _ := ioutil.ReadAll(data)

	contentType := http.DetectContentType(fileBytes)

	// check content type include in allowed content type
	if !IsAllowedContentType(contentType) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Content type is not allowed"})
		return
	}

	// file extension
	ext := path.Ext(file.Filename)

	// generate unique file name
	fileId := fmt.Sprintf("%d%s", time.Now().Unix(), ext)

	// file path
	destination := path.Join(UPLOAD_DESTINATION, fileId)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, destination)

	// redirect to /gallery
	c.Redirect(http.StatusMovedPermanently, "/gallery")
}

func GalleryRoute(c *gin.Context) {
	// get all file in upload dir
	files, _ := ioutil.ReadDir(UPLOAD_DESTINATION)

	// make image url
	images := make([]string, 0)
	for _, file := range files {
		// is file is directory
		if file.IsDir() {
			continue
		}
		images = append(images, "/upload/"+file.Name())
	}

	// render template
	c.HTML(http.StatusOK, "gallery.html", gin.H{
		"Images": images,
	})
}
