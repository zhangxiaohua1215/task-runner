package controller

import "task-runner/service"

type ApiGroup struct {
	Script
	Task
}

var AppApiGroup = new(ApiGroup)

var appService = service.AppServiceGroup