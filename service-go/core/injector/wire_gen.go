// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"magician/api"
	"magician/core/container"
	"magician/core/provider"
	"magician/model/dao"
	"magician/router"
	"magician/service"
)

// Injectors from wire.go:

func BuildContainer() (*container.Container, func(), error) {
	db, cleanup, err := provider.BuildGorm()
	if err != nil {
		return nil, nil, err
	}
	user := &dao.User{
		DB: db,
	}
	test := &service.Test{
		UserDao: user,
	}
	apiTest := &api.Test{
		Test: test,
	}
	routerRouter := &router.Router{
		Test: apiTest,
	}
	engine := provider.BuildHTTPHandler(routerRouter)
	containerContainer := &container.Container{
		Engine: engine,
	}
	return containerContainer, func() {
		cleanup()
	}, nil
}
