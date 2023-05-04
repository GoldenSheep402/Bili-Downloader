package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"time"
)

type Cookie struct {
	_uuid string
	buvid string
}

type GetQrcodeRep struct {
	Code   int  `json:"code"`
	Status bool `json:"status"`
	Ts     int  `json:"ts"`
	Data   struct {
		URL      string `json:"url"`
		OAuthKey string `json:"oauthKey"`
	} `json:"data"`
}

func GenCookie(c *Cookie) *Cookie {
	uid := uuid.New()
	uuidString := uid.String()
	result := fmt.Sprintf("%sinfoc", uuidString[:8]+"-"+uuidString[8:12]+"-"+uuidString[12:16]+"-"+uuidString[16:20]+"-"+uuidString[20:])
	c._uuid = result

	buvid := uuid.New()
	buvidString := buvid.String()
	result = fmt.Sprintf("%sinfoc", buvidString[:8]+"-"+buvidString[8:12]+"-"+buvidString[12:16]+"-"+buvidString[16:20]+"-"+buvidString[20:])
	c.buvid = result

	return c
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
		"Cookie":     []string{fmt.Sprintf("_uuid=%s; buvid=%s", c._uuid, c.buvid)},
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

	return data
}

// TODO
func ListenQrcode(in *GetQrcodeRep) {

}
