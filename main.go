package main

import (
	"task-runner/job"
	"task-runner/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RegisterRouter(r)
	gin.ForceConsoleColor()
	job.RunWorker(3)
	r.Run(":8080")
}
