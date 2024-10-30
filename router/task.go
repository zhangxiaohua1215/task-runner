package router

import (

	"github.com/gin-gonic/gin"
)

type TaskRouter struct {}

func (r *TaskRouter) RegisterTaskRouter(e *gin.Engine) {
	api := e.Group("/task")
	{
		api.POST("/execute", appController.ExecuteTask)
	}
}
