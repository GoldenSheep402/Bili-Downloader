package merger

import (
	"fmt"
	"os"
	"os/exec"
)

func MergeVideoAndAudio(videoPath string, audioPath string, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-y", "-i", videoPath, "-i", audioPath, "-c", "copy", outputPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Please install ffmpeg in your system")
		os.Exit(1)
		return err
	}
	return nil
}
