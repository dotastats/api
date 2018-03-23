package handler

import (
	"dotaapi/config"
	"dotamaster/middleware"
	"dotamaster/utilities/ulog"

	"github.com/gin-gonic/gin"
)

func InitEngine(conf *config.Config) *gin.Engine {
	if conf.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware(conf.App.Whitelist))
	r.Use(gin.LoggerWithWriter(ulog.Logger().Request))

	if conf.App.Debug {
		r.Use(gin.Logger())
	}

	// Ping
	pingHandler := pingHandler{}
	groupPing := r.Group("/")
	{
		GET(groupPing, "", pingHandler.Ping)
	}
	return r
}

func GET(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "GET", relativePath, f)
}

func POST(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "POST", relativePath, f)
}

func PUT(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "PUT", relativePath, f)
}

func DELETE(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "DELETE", relativePath, f)
}

func route(group *gin.RouterGroup, method string, relativePath string, f func(*gin.Context)) {
	hanld := middleware.ErrorHandler(group.BasePath() + relativePath)
	switch method {
	case "POST":
		group.POST(relativePath, hanld, f)
	case "GET":
		group.GET(relativePath, hanld, f)
	case "PUT":
		group.PUT(relativePath, hanld, f)
	case "DELETE":
		group.DELETE(relativePath, hanld, f)
	}
}
