package handler

import "github.com/gin-gonic/gin"

type pingHandler struct {
}

func (pingHandler) Ping(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "pong"})
}
