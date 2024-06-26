// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"foodV5/common/middleware"
	"foodV5/common/pkg/wechat"
	"foodV5/common/repo"
	"foodV5/common/service"
	"foodV5/foodV5-api/app"
	"foodV5/foodV5-api/app/api"
	"foodV5/foodV5-api/app/router"
	service2 "foodV5/foodV5-api/app/service"
)

// Injectors from wire.go:

func wireApp() (*app.Application, error) {
	engine := router.NewGin()
	client := app.NewRedis()
	middlewareMiddleware := &middleware.Middleware{
		Client: client,
	}
	miniProgram := wechat.NewWechatMini()
	db, err := app.NewGorm()
	if err != nil {
		return nil, err
	}
	repoRepo := &repo.Repo{}
	userRepo := &repo.UserRepo{
		DB:   db,
		Repo: repoRepo,
	}
	configRepo := &repo.ConfigRepo{
		DB: db,
	}
	inviteRepo := &repo.InviteRepo{
		DB:   db,
		Repo: repoRepo,
	}
	inviteService := &service.InviteService{
		InviteRepo: inviteRepo,
	}
	integralRepo := &repo.IntegralRepo{
		DB:   db,
		Repo: repoRepo,
	}
	accountService := &service2.AccountService{
		Mini:          miniProgram,
		Redis:         client,
		DB:            db,
		UserRepo:      userRepo,
		ConfigRepo:    configRepo,
		InviteRepo:    inviteRepo,
		InviteService: inviteService,
		IntegralRepo:  integralRepo,
	}
	accountApi := &api.AccountApi{
		AccountService: accountService,
	}
	routerRouter := &router.Router{
		Engine:     engine,
		Middleware: middlewareMiddleware,
		AccountApi: accountApi,
	}
	application := &app.Application{
		Router: routerRouter,
	}
	return application, nil
}
