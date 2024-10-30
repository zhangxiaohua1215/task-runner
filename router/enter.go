package router

import "task-runner/controller"

type RouterGroup struct {
	ScriptRouter
	TaskRouter
}

var AppRouterGroup = new(RouterGroup)

var appController = controller.AppApiGroup