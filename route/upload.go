package route

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wuttinanhi/go-web-file-upload/config"
	"github.com/wuttinanhi/go-web-file-upload/util"
)

func UploadRoute(c *gin.Context) {
	// Single file
	file, _ := c.FormFile("file")

	// read file as byte
	data, _ := file.Open()
	defer data.Close()
	fileBytes, _ := ioutil.ReadAll(data)

	contentType := http.DetectContentType(fileBytes)

	// check content type include in allowed content type
	if !util.IsAllowedContentType(contentType) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Content type is not allowed"})
		return
	}

	// file extension
	ext := path.Ext(file.Filename)

	// generate unique file name
	fileId := fmt.Sprintf("%d%s", time.Now().Unix(), ext)

	// file path
	destination := path.Join(config.GetConfig().UPLOAD_DESTINATION, fileId)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, destination)

	// redirect to /gallery
	c.Redirect(http.StatusMovedPermanently, "/gallery")
}
