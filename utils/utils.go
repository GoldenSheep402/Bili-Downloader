package utils

import (
	"fmt"
	"os"
	"regexp"
)

func CheckBid(bid string) bool {
	re := regexp.MustCompile(`^BV1?[a-zA-Z0-9]{10}$`)
	if re.MatchString(bid) {
		return true
	} else {
		return false
	}
}

// TODO
func GetBidByUrl() string {
	var url string
	fmt.Scan(&url)
	print(url)

	// regular match
	re := regexp.MustCompile(`(BV|BV1|BV[a-zA-Z0-9]{2})[a-zA-Z0-9]+`)
	match := re.FindStringSubmatch(url)
	return match[0]
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
