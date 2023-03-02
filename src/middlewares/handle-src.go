package middlewares

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gfile"
	"s3cache/src/utils"
	"s3cache/src/utils/storage"
	"strings"
)

func HandleSrc(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		srcPath := strings.Replace(c.Request.URL.Path, "/file/", "", 1)
		// 从远程下载文件
		bucket := c.DefaultQuery("bucket", "")
		localSrcPath := ""
		if bucket != "" {
			localSrcPath = utils.DataPath("src", bucket, srcPath)
		} else {
			// 作为本地使用时
			localSrcPath = utils.DataPath("src", srcPath)
		}
		if utils.IsExist(localSrcPath) == false {
			// 将下载过的数据放入redis中
			if utils.GetConfig().Server.RemoteEnable == false {
				return
			}
			if bucket == "" {
				return
			}
			err := storage.DownloadObject(bucket, srcPath, utils.DataPath("src", bucket, srcPath))
			if err != nil {
				return
			}
		}
		baseName := gfile.Basename(srcPath)
		baseNameInfo := strings.Split(baseName, ".")
		if len(baseNameInfo) > 1 && utils.IsImage(baseNameInfo[1]) {
			format := c.DefaultQuery("f", "")
			width := c.DefaultQuery("w", "")
			height := c.DefaultQuery("h", "")
			mode := c.DefaultQuery("m", "")
			if format != "" || width != "" || height != "" {
				// 重定向到缓存文件
				unionName := baseName + "?w=" + width + "&h=" + height + "&f=" + format + "&m=" + mode
				hash := md5.Sum([]byte(unionName))
				hashStr := hex.EncodeToString(hash[:])
				c.Request.URL.Path = "/cache/" + hashStr[0:2] + "/" + hashStr[2:4] + "/" + hashStr + "." + format
				c.Request.URL.RawQuery = "src=" + srcPath + "&" + c.Request.URL.RawQuery
				r.HandleContext(c)
			}
		} else {
			if bucket != "" {
				c.Request.URL.Path = "/src/" + bucket + "/" + srcPath
			} else {
				c.Request.URL.Path = "/src/" + srcPath
			}
			c.Request.URL.RawQuery = ""
			r.HandleContext(c)
		}
	}
}
