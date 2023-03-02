package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/duke-git/lancet/fileutil"
	"github.com/xiebinbin/m3u8"
	"net/http"
	"os"
	"s3cache/src/utils"
	"s3cache/src/utils/storage"
)

func buildM3U8(src string) *bytes.Buffer {
	nodes := utils.GetConfig().Server.Nodes
	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	p, _, err := m3u8.DecodeFrom(bufio.NewReader(f), true)
	if err != nil {
		panic(err)
	}
	playFile := p.(*m3u8.MediaPlaylist)
	for i, segment := range playFile.Segments {
		if segment == nil {
			break
		}
		nodeIndex := uint8(i) % uint8(len(nodes))
		node := nodes[nodeIndex]
		playFile.Segments[i].URI = node + "file/" + segment.URI + "?bucket=hls-ts"
	}
	return playFile.Encode()
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	bucket := "hls-m3u8"
	objectName := r.URL.Path
	localSrcPath := utils.DataPath("src", bucket, objectName)
	fmt.Println("path", localSrcPath)
	fmt.Println("scheme", r.URL.Scheme)
	if fileutil.IsExist(localSrcPath) == false {
		err := storage.DownloadObject(bucket, objectName, localSrcPath)
		if err != nil {
			return
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	_, err := w.Write(buildM3U8(localSrcPath).Bytes())
	if err != nil {
		return
	}
}
func main() {
	http.HandleFunc("/", IndexHandler)
	err := http.ListenAndServe(utils.GetConfig().Server.M3u8Address, nil)
	if err != nil {
		return
	}
}
