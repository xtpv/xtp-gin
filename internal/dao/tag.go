package dao

import (
	"github.com/xtpv/xtp-gin/internal/model"
)

func (d Dao) GetTagList(state uint8) (tags []*model.Tag, err error) {
	err = d.engine.Model(model.Tag{}).Where("state = ?", state).Find(&tags).Error
	return
}
