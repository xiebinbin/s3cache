package storage

import (
	"path"
	"s3cache/src/utils"
)

func initPath() {
	dataPath := utils.GetConfig().Server.DataPath
	if dataPath == "" {
		dataPath = utils.AppPath("data")
	}
	srcPath := path.Join(dataPath, "src")
	cachePath := path.Join(dataPath, "cache")
	utils.MakeDir(srcPath)
	utils.MakeDir(cachePath)
}
func init() {
	initPath()
	initMinioClient()
}
