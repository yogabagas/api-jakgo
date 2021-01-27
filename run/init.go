package command

import (
	"my-github/api-jakgo/controller"
	"my-github/api-jakgo/service"
	"my-github/api-jakgo/usecase"
)

func NewJakGoRegistry() service.JakGoService {

	urlRsu := "http://api.jakarta.go.id/v1/rumahsakitumum"
	urlKel := "http://api.jakarta.go.id/v1/kelurahan"
	urls := []string{urlRsu, urlKel}

	return usecase.NewJakGoService(
		"LdT23Q9rv8g9bVf8v/fQYsyIcuD14svaYL6Bi8f9uGhLBVlHA3ybTFjjqe+cQO8k",
		urls...,
	)
}

func NewJakGoController() controller.JakGoController {
	return controller.NewJakGoController(NewJakGoRegistry())
}
