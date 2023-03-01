package utils

import (
	"github.com/wxnacy/wgo/arrays"
	"io"
	"log"
	"os"
	"path"
)

func IsImage(ext string) bool {
	includes := []string{"jpg", "jpeg", "png", "webp", "bmp", "tiff"}
	index := arrays.ContainsString(includes, ext)
	if index == -1 {
		return false
	}
	return true
}
func AppPath(elem ...string) string {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal("路径获取失败!¬")
	}
	return path.Join(rootPath, path.Join(elem...))
}

func DataPath(elem ...string) string {
	dataPath := GetConfig().Server.DataPath
	if dataPath == "" {
		return AppPath("data", path.Join(elem...))
	}
	return path.Join(dataPath, path.Join(elem...))
}
func MakeDir(dir string) {
	if IsExist(dir) == false {
		err := os.MkdirAll(dir, 0776)
		if err != nil {
			log.Fatal("文件夹" + dir + "创建失败")
		}
	}
}
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
func ReadFile(path string) []byte {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal("config file err")
	}
	buffer := make([]byte, fileInfo.Size())
	fileHandler, err := os.Open(path)
	if err != nil {
		log.Fatal("config file err")
	}
	_, err = io.ReadFull(fileHandler, buffer)
	if err != nil {
		log.Fatal("config file err")
	}
	return buffer
}
