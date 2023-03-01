package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gfile"
	"log"
	"os"
	"path/filepath"
	imagepro "s3cache/src/image-pro"
	"s3cache/src/utils"
	"strconv"
	"strings"
)

func HandleCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		cachePath := utils.DataPath(c.Request.URL.Path)
		if utils.IsExist(cachePath) == false {
			baseName := gfile.Basename(cachePath)
			baseNameInfo := strings.Split(baseName, ".")
			if len(baseNameInfo) > 1 && utils.IsImage(baseNameInfo[1]) {
				log.Println("处理图片")
				src := c.DefaultQuery("src", "")
				bucket := c.DefaultQuery("bucket", "")
				srcPath := utils.DataPath("src", bucket, src)
				format := c.DefaultQuery("f", "")
				width, err := strconv.Atoi(c.DefaultQuery("w", "0"))
				if err != nil {
					return
				}
				height, err := strconv.Atoi(c.DefaultQuery("h", "0"))
				if err != nil {
					return
				}
				mode := c.DefaultQuery("m", "")
				buff, err := imagepro.New(srcPath, &imagepro.Options{
					Format:  format,
					Mode:    mode,
					Width:   width,
					Height:  height,
					Quality: 85,
				}).Transform().Encode()
				if err != nil {
					return
				}
				if err != nil {
					log.Println(err)
					return
				}
				utils.MakeDir(filepath.Dir(cachePath))

				// 保存处理后文件
				f, err := os.Create(cachePath)
				if err != nil {
					log.Println(err)
					return
				}
				_, err = f.Write(buff)
				if err != nil {
					log.Println(err)
					return
				}
				err = f.Close()
				if err != nil {
					log.Println(err)
					return
				}
			}
		}

	}
}
