package service

import (
	"context"
	"my-github/api-jakgo/domain/model"
)

type JakGoService interface {
	JakGoRSU(ctx context.Context) (rsu *model.RSUResponse, err error)
	JakGoKelurahan(ctx context.Context) (kel *model.KelurahanResponse, err error)
}
