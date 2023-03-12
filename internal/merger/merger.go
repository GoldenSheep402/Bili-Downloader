package merger

import "C"

import (
	"fmt"
	"os/exec"
)

func MergeVideoAndAudio(videoPath string, audioPath string, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-y", "-i", videoPath, "-i", audioPath, "-c", "copy", outputPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to merge video and audio:", err)
		return err
	}
	fmt.Println("Merge video and audio successfully.")
	return nil
}
