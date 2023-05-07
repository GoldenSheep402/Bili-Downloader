package main

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"time"
)

type Cookie struct {
	_uuid          string
	buvid3         string
	fingerprint    string
	sid            string
	buvid_fp       string
	buvid_fp_plain string
}

type GetQrcodeRep struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Ts     int    `json:"ts"`
	Data   struct {
		URL      string `json:"url"`
		OAuthKey string `json:"oauthKey"`
	} `json:"data"`
}

type Response struct {
	Code   int   `json:"code"`
	Status bool  `json:"status"`
	Ts     int64 `json:"ts"`
	Data   struct {
		URL string `json:"url"`
	} `json:"data"`
}

func RandomString(length int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)

	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(err)
		}
		result[i] = letters[n.Int64()]
	}

	return string(result)
}

func GenCookie(c *Cookie) {
	randString := RandomString(32)
	result := fmt.Sprintf("%sinfoc", randString[:8]+"-"+randString[8:12]+"-"+randString[12:16]+"-"+randString[16:20]+"-"+randString[20:])
	c._uuid = result

	randString = RandomString(32)
	result = fmt.Sprintf("%sinfoc", randString[:8]+"-"+randString[8:12]+"-"+randString[12:16]+"-"+randString[16:20]+"-"+randString[20:])
	c.buvid3 = result

	c.fingerprint = RandomString(32)
	c.sid = RandomString(8)

	randString = RandomString(32)
	result = fmt.Sprintf("%sinfoc", randString[:8]+"-"+randString[8:12]+"-"+randString[12:16]+"-"+randString[16:20]+"-"+randString[20:])
	c.buvid_fp = result

	result = fmt.Sprintf(result)
	c.buvid_fp_plain = result
}

func GetQrcode(c *Cookie) *GetQrcodeRep {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest("GET", "https://passport.bilibili.com/qrcode/getLoginUrl", nil)
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Referer":    []string{"https://passport.bilibili.com/login?from_spm_id=333.851.top_bar.login_window"},
		"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.51 Safari/537.36 Edg/93.0.961.27"},
		"Cookie":     []string{fmt.Sprintf("_uuid=%s; buvid=%s", c._uuid, c.buvid3)},
	}

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
	var data *GetQrcodeRep
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	url := data.Data.URL

	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		fmt.Println("无法生成二维码：", err)
		return nil
	}
	// 打印二维码到控制台
	fmt.Println(qr.ToSmallString(false))

	return data
}

// TODO
func ListenQrcode(c *Cookie, getQrcodeRep *GetQrcodeRep) {
	client := &http.Client{
		Timeout: time.Second * 180,
	}

	data := url.Values{}
	fmt.Println(getQrcodeRep.Data.OAuthKey)
	data.Set("oauthKey", getQrcodeRep.Data.OAuthKey)
	payload := bytes.NewBufferString(data.Encode())

	req, err := http.NewRequest("POST", "https://passport.bilibili.com/qrcode/getLoginInfo", payload)
	if err != nil {
		panic(err)
	}

	fmt.Println("===")
	fmt.Println(c._uuid)
	fmt.Println(c.sid)
	fmt.Println(c.buvid_fp)
	fmt.Println(c.buvid_fp_plain)
	fmt.Println(c.fingerprint)
	fmt.Println(c.buvid3)
	fmt.Println("=)==")
	req.Header = http.Header{
		"Referer": []string{"https://passport.bilibili.com/login?from_spm_id=333.851.top_bar.login_window"},
		//"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.51 Safari/537.36 Edg/93.0.961.27"},
	}
	req.Header.Add("Cookie", fmt.Sprintf("_uuid=%s; buvid3=%s; fingerprint=%s; sid=%s; buvid_fp=%s; buvid_fp_plain=%s", c._uuid, c.buvid3, c.fingerprint, c.sid, c.buvid_fp, c.buvid_fp_plain))
	fmt.Println(req.Header)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var obj map[string]interface{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println("翻译后的JSON对象：")
	fmt.Printf("状态: %v\n", obj["status"].(bool))
	fmt.Printf("数据: %v\n", obj["data"].(float64))
	fmt.Printf("消息: %s\n", obj["message"].(string))
}

func main() {
	c := &Cookie{}
	GenCookie(c)

	resp := GetQrcode(c)
	fmt.Println(resp)
	time.Sleep(time.Second * 6)
	ListenQrcode(c, resp)
	return
}
