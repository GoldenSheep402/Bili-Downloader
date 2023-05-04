package downloader

import (
	"Bili-Downloader/define"
	"Bili-Downloader/internal/merger"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Download(VideoInfo *define.VideoInfo) {
	//SaveData(VideoInfo)
	size := len(VideoInfo.VideoUrl)
	for i := 0; i < size; i++ {
		// 下载视频
		DownloadVideo(VideoInfo, i)
		// 下载音频
		DownloadAudio(VideoInfo, i)
		// 合并音视频
		merger.MergeVideoAndAudio(
			"./download_path/video.mp4",
			"./download_path/audio.mp4",
			"./download_path/"+VideoInfo.Bvid+"["+VideoInfo.Part[i]+"].mp4")

		err := os.Remove("./download_path/video.mp4")
		if err != nil {
			// 处理删除文件时出现的错误
		}

		err = os.Remove("./download_path/audio.mp4")
		if err != nil {
			// 处理删除文件时出现的错误
		}

	}
}

func SaveData(VideoInfo *define.VideoInfo) {
	// Convert the struct to JSON bytes
	jsonBytes, _ := json.MarshalIndent(VideoInfo, "", "    ")

	// Write the JSON bytes to a file
	filename := "data.json"
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	ioutil.WriteFile(filepath.Join(dir, filename), jsonBytes, 0644)
}
