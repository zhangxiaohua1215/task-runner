package router

import "github.com/gin-gonic/gin"

type ScriptRouter struct{}

func (r *ScriptRouter) RegisterScriptRouter(e *gin.Engine) {
	api := e.Group("/script")
	{
		api.POST("/upload", appController.UploadScript)
		api.POST("/list", appController.ListScript)
		api.POST("/detail", appController.DetailScript)
		api.GET("/download/:id", appController.DownloadScript)
	}
}
