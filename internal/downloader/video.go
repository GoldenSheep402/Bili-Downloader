package downloader

import (
	"Bili-Downloader/define"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
)

func DownloadVideo(VideoInfo *define.VideoInfo, index int) {

	// 创建http请求客户端
	client := &http.Client{}

	// 创建http请求
	req, err := http.NewRequest("GET", VideoInfo.VideoUrl[index], nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置请求头部信息
	req.Header.Set("referer", "https://www.bilibili.com/")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0")

	// 发送http请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 创建文件
	out, err := os.Create("./download_path/video.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	// 获取响应的body大小
	size := resp.ContentLength

	// 创建进度条
	bar := progressbar.NewOptions64(size, progressbar.OptionSetRenderBlankState(true))

	// 下载文件
	buf := make([]byte, 1024)
	_, err = io.CopyBuffer(io.MultiWriter(out, bar), resp.Body, buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 下载完成
	fmt.Println("Video:" + "[" + VideoInfo.Part[index] + "]" + "下载完成！")
}
