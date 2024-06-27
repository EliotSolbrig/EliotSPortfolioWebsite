package router

import (
    "main/service"
)

var basePasefiles []string = []string{
    "templates/base.html",
    "templates/header.html",
}

type Router struct {
    Service service.IService
}

func NewRouter(service service.IService) *Router {
    return &Router{
        Service: service,
    }
}
