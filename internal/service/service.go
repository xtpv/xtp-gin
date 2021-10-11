package service

import (
	"context"
	G "github.com/xtpv/xtp-gin/global"
	"github.com/xtpv/xtp-gin/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(G.DB)
	return svc
}
