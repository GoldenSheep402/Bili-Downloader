package main

import (
	"Bilibili-DL/define"
	"Bilibili-DL/internal/bili"
	"Bilibili-DL/internal/downloader"
	"Bilibili-DL/internal/merger"
	"fmt"
)

func main() {

	var bvid string
	fmt.Print("Please enter the BID: ")
	fmt.Scanln(&bvid)

	VideoInfo := define.VideoInfo{
		Bvid: bvid,
	}

	// 获取视频api链接
	err := bili.GetApi(&VideoInfo)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// 获取实际下载链接
	bili.GetUrl(&VideoInfo)

	// 下载音视频
	downloader.DownloadVideo(&VideoInfo)
	downloader.DownloadAudio(&VideoInfo)

	// 合并音视频
	merger.MergeVideoAndAudio("./download_path/video.mp4", "./download_path/audio.mp4", "./"+VideoInfo.Bvid+".mp4")

	fmt.Println("Download completed successfully!")
	fmt.Println("Press any key to exit...")
	fmt.Scanln()
}
