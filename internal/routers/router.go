package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Pong)
		v1.POST("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.DELETE("/ping", Pong)

	}
	return r

}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "namdang-fdp")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...abc" + name,
		"uid":     uid,
		"users":   []string{"cr7", "namdang-fdp"},
	})
}
