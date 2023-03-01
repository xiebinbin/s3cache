package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"s3cache/src/middlewares"
	"s3cache/src/utils"
)

func RunServer() {
	log.Println("run server " + utils.GetConfig().Server.Address)
	if utils.GetConfig().Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	var r = gin.Default()

	// 第一层用户接收用户请求
	fileRoute := r.Group("/file")
	fileRoute.Use(middlewares.HandleSrc(r))
	fileRoute.Static("/", utils.DataPath("src"))
	// 用于跳转源文件
	srcRoute := r.Group("/src")
	srcRoute.Static("/", utils.DataPath("src"))
	// 用于跳转已处理的文件
	cacheRoute := r.Group("/cache")
	cacheRoute.Use(middlewares.HandleCache())
	cacheRoute.Static("/", utils.DataPath("cache"))

	err := r.Run(utils.GetConfig().Server.Address)
	if err != nil {
		return
	}
}
func main() {
	RunServer()
}
