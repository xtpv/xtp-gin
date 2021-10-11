package service

import "github.com/xtpv/xtp-gin/internal/model"

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

func (s Service) GetTagList(req TagListRequest) ([]*model.Tag, error) {
	return s.dao.GetTagList(req.State)
}
