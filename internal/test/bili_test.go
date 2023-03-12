package test

import (
	"Bilibili-DL/define"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type Response struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Ttl     int            `json:"ttl"`
	Data    []ResponseData `json:"data"`
}

type ResponseData struct {
	Cid       int    `json:"cid"`
	Page      int    `json:"page"`
	From      string `json:"from"`
	Part      string `json:"part"`
	Duration  int    `json:"duration"`
	Vid       string `json:"vid"`
	Weblink   string `json:"weblink"`
	Dimension struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Rotate int `json:"rotate"`
	} `json:"dimension"`
}

func GetVideoUrl(videoInfo *define.VideoInfo) {
	url := define.BaseUrlCid + videoInfo.Bvid

	// Create an HTTP client with headers
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Second * 30,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	//req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36")

	// Send the GET request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON data
	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	// Get the cid
	videoInfo.Cid = strconv.Itoa(data.Data[0].Cid)
	videoInfo.Url = define.BaseUrlVideo + "&cid=" + videoInfo.Cid + "&bvid=" + videoInfo.Bvid
	//fmt.Fprint(os.Stdout, "Url: "+videoInfo.Url+"\n")

}

func TestGetVideoUrl(t *testing.T) {
	testCases := []struct {
		Bvid        string
		ExpectedCid string
		ExpectedUrl string
	}{
		{
			Bvid:        "BV1JN411F7k8",
			ExpectedCid: "1048988381",
			ExpectedUrl: "https://api.bilibili.com/x/player/playurl?fnval=80&cid=1048988381&bvid=BV1JN411F7k8",
		},
		{
			Bvid:        "BV1p14y1c7o8",
			ExpectedCid: "1003897861",
			ExpectedUrl: "https://api.bilibili.com/x/player/playurl?fnval=80&cid=1003897861&bvid=BV1p14y1c7o8",
		},
	}

	for _, tc := range testCases {
		videoInfo := &define.VideoInfo{
			Bvid: tc.Bvid,
		}
		GetVideoUrl(videoInfo)

		if videoInfo.Cid != tc.ExpectedCid {
			t.Errorf("GetVideoUrl(%v): got Cid = %v, want %v", videoInfo, videoInfo.Cid, tc.ExpectedCid)
		}
		if videoInfo.Url != tc.ExpectedUrl {
			t.Errorf("GetVideoUrl(%v): got Url = %q, want %q", videoInfo, videoInfo.Url, tc.ExpectedUrl)
		}
	}
}
