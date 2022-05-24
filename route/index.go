package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
