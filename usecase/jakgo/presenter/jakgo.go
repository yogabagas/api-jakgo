package presenter

import (
	"context"
	"my-github/api-jakgo/domain/model"
)

type JakGo interface {
	ResponseJakGo(context.Context, []*model.JakGo) ([]*model.JakGo, error)
}
