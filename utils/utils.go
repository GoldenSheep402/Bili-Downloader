package utils

import (
	"fmt"
	"os"
)

// TODO
func GetBidByUrl() {

}

func PrintInfo() {
	fmt.Print("[Bilibili Downloader]\n" +
		"By: A Sheep made of gold\n\n")
}

func PrintQuitInfo() {
	fmt.Print("Download completed successfully!\n" +
		"Thank u for using it.\n" +
		"Have a nice day!\n" +
		"Press any key to exit...\n")
	fmt.Scanln()
	os.Exit(1)
}
