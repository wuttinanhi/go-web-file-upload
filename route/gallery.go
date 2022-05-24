package route

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuttinanhi/go-web-file-upload/config"
)

func GalleryRoute(c *gin.Context) {
	// get all file in upload dir
	files, _ := ioutil.ReadDir(config.GetConfig().UPLOAD_DESTINATION)

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
