package service

type ServiceGroup struct {
	Task
	Script
}

var AppServiceGroup = new(ServiceGroup)
