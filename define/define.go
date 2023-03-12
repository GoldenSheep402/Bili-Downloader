package define

type VideoInfo struct {
	Bvid     string
	Cid      string
	Url      string
	VideoUrl []string
	AudioUrl []string
}

type VideoResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    Data   `json:"data"`
}

type Data struct {
	//AcceptDescription []string `json:"accept_description"`
	Dash Dash `json:"dash"`
}

type Dash struct {
	Video []Video `json:"video"`
	Audio []Audio `json:"audio"`
}

type Video struct {
	Id      int    `json:"id"`
	BaseUrl string `json:"base_url"`
}

type Audio struct {
	Id      int    `json:"id"`
	BaseUrl string `json:"base_url"`
}

const BaseUrlCid = "https://api.bilibili.com/x/player/pagelist?bvid="
const BaseUrlVideo = "https://api.bilibili.com/x/player/playurl?fnval=80"
