package handler

import (
	"dotaapi/app/entity"

	"github.com/gin-gonic/gin"
)

type teamHandler struct {
	teamEntity entity.TeamEntity
}

func (teamHandler) GetInfo(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "pong"})
}

func (teamHandler) GetMatches(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "pong"})
}

func (teamHandler) GetHistory(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "pong"})
}

func (teamHandler) GetF10kMatches(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "pong"})
}

func (teamHandler) GetFBMatches(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "pong"})
}
