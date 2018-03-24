package handler

import (
	"dotaapi/app/entity"

	"github.com/gin-gonic/gin"
)

type matchHandler struct {
	matchEntity entity.MatchEntity
}

func (matchHandler) GetList(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "pong"})
}

func (matchHandler) GetDetail(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "pong"})
}
