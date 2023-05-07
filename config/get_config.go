package config

import (
	"Bili-Downloader/internal/bili"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	BID      string `json:"BID"`
	SESSDATA string `json:"SESSDATA"`
}

func GetConfig() (string, string, error) {
	// 获取配置文件
	configFile, err := os.Open("config/config.json")
	if err != nil {
		return "", "", fmt.Errorf("failed to open config file: %s", err)
	}
	defer configFile.Close()

	// 读取配置文件
	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return "", "", fmt.Errorf("failed to read config file: %s", err)
	}

	// 解析配置文件
	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse config file: %s", err)
	}

	// 获取SESSDATA
	var sessdata string
	fmt.Print("Do you want to get SESSDATA by qrCode? (yes/no) ")
	var answer string
	_, err = fmt.Scanln(&answer)
	if err != nil {
		return "", "", fmt.Errorf("failed to read user input: %s", err)
	}
	if strings.ToLower(answer) == "yes" || strings.ToLower(answer) == "y" {
		sessdata = bili.GetSESSDATA()
		config.SESSDATA = sessdata
		// 更新配置文件
		configBytes, err = json.MarshalIndent(config, "", "    ")
		if err != nil {
			return "", "", fmt.Errorf("failed to update config file: %s", err)
		}
		err = ioutil.WriteFile("config/config.json", configBytes, 0644)
		if err != nil {
			return "", "", fmt.Errorf("failed to update config file: %s", err)
		}
	} else {
		sessdata = config.SESSDATA
	}

	// 获取BID
	bid := config.BID

	return bid, sessdata, nil
}
