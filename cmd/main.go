package main

import (
	"Bili-Downloader/config"
	"Bili-Downloader/define"
	"Bili-Downloader/internal/bili"
	"Bili-Downloader/internal/downloader"
	"fmt"
)

func main() {

	bid, sessdata, err := config.GetConfig()
	if err != nil {
		fmt.Println("error: ", err)
		fmt.Println("Press any key to exit...")
		fmt.Scanln()
		return
	}

	VideoInfo := define.VideoInfo{
		Bvid: bid,
	}

	// 获取视频api链接
	err = bili.GetApi(&VideoInfo)
	if err != nil {
		fmt.Println("error: ", err)
		fmt.Println("Press any key to exit...")
		fmt.Scanln()
		return
	}

	// 获取实际下载链接
	bili.GetUrl(&VideoInfo, sessdata)

	// 下载音视频并合并
	downloader.Download(&VideoInfo)

	fmt.Println("Download completed successfully!")
	fmt.Println("Press any key to exit...")
	fmt.Scanln()
}
