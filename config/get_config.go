package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

	// 获取BID和SESSDATA
	bid := config.BID
	sessdata := config.SESSDATA

	return bid, sessdata, nil
}
