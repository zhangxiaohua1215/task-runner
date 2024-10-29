package router

import (
	"task-runner/handler"

	"github.com/gin-gonic/gin"
)

// 注册路由
func RegisterRouter(r *gin.Engine) {
	
	s := new(handler.ScriptHandler)
	r.POST("/upload", s.UploadScript)

	t := new(handler.TaskHandler)
	r.POST("/execute", t.ExecuteTask)
}
