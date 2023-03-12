package main

import (
	"Bilibili-DL/config"
	"Bilibili-DL/define"
	"Bilibili-DL/internal/bili"
	"Bilibili-DL/internal/downloader"
	"Bilibili-DL/internal/merger"
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

	// 下载音视频
	downloader.DownloadVideo(&VideoInfo)
	downloader.DownloadAudio(&VideoInfo)

	// 合并音视频
	merger.MergeVideoAndAudio("./download_path/video.mp4", "./download_path/audio.mp4", "./"+VideoInfo.Bvid+".mp4")

	fmt.Println("Download completed successfully!")
	fmt.Println("Press any key to exit...")
	fmt.Scanln()
}
