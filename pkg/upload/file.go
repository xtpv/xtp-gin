package upload

import (
	G "github.com/xtpv/xtp-gin/global"
	"github.com/xtpv/xtp-gin/pkg/utils"
	"path"
	"time"
)

type FileType int

const (
	FileImage FileType = iota + 1
	FileExcel
	FileTxt
)

func GetFileName(name string) string {
	ext := path.Ext(name)
	fileName := utils.EncodeMd5(name + time.Now().String())

	return fileName + ext
}

func GetSavePath() string {
	return G.AppConfig.UploadSavePath
}
