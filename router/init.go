package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	r := gin.Default()
	AppRouterGroup.RegisterScriptRouter(r)
	AppRouterGroup.RegisterTaskRouter(r)
	return r

}