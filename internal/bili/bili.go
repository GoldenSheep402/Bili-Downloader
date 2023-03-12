package bili

import (
	"Bili-Downloader/define"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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

func GetApi(videoInfo *define.VideoInfo) error {
	if videoInfo.Bvid == "" {
		err := errors.New("BV ID is required")
		return err
	}

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
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse the JSON data
	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	// Get the cid
	videoInfo.Part = make([]string, len(data.Data))
	videoInfo.Cid = make([]string, len(data.Data))
	videoInfo.Url = make([]string, len(data.Data))
	for i := 0; i < len(data.Data); i++ {
		print(data.Data[i].Part + "\n")
		videoInfo.Part[i] = data.Data[i].Part
		videoInfo.Cid[i] = strconv.Itoa(data.Data[i].Cid)
		videoInfo.Url[i] = define.BaseUrlVideo + "&cid=" + videoInfo.Cid[i] + "&bvid=" + videoInfo.Bvid
	}

	err = os.MkdirAll("./download_path/data", os.ModePerm)
	if err != nil {
		panic(err)
	}
	// Save the JSON data to a file
	file, err := os.Create("./download_path/data/cid.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.Write(body); err != nil {
		return err
	}
	return nil
}

func GetUrl(videoInfo *define.VideoInfo, SESSDATA string) {
	// create an HTTP client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Second * 30,
	}

	// make a slice to store the video and audio urls
	videoInfo.VideoUrl = make([]string, len(videoInfo.Cid))
	videoInfo.AudioUrl = make([]string, len(videoInfo.Cid))
	for index := 0; index < len(videoInfo.Cid); index++ {
		// create an HTTP request
		req, err := http.NewRequest("GET", videoInfo.Url[index], nil)
		if err != nil {
			panic(err)
		}

		cookie := &http.Cookie{
			Name:   "SESSDATA",
			Value:  SESSDATA,
			Domain: "example.com",
			Path:   "/",
		}
		req.AddCookie(cookie)

		// headers
		req.Header.Set("referer", "https://www.bilibili.com/")
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0")

		// send the request
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// read
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// parse the JSON data
		var data define.VideoResponse
		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}

		// Get video and audio data and titile(part)
		videoInfo.VideoUrl[index] = data.Data.Dash.Video[0].BaseUrl
		videoInfo.AudioUrl[index] = data.Data.Dash.Audio[0].BaseUrl
	}
	fmt.Fprint(os.Stdout, "Get video and audio data successfully!\n")
}
