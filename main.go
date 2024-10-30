package main

import (
	"task-runner/job"
	"task-runner/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	r := router.Init()
	job.RunWorker(3)
	r.Run(":8080")
}
