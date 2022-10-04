package rest

import (
	_ "github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Eng *gin.Engine

func init() {
	Eng = gin.Default()

	Eng.GET("/api/v1/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": http.StatusOK,
			"msg":  "hello",
		})
	})
}
