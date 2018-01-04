package petglobal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"petApi/models"
	"time"
)

//AccessToken 微信AccessToken
var AccessToken = ""

func GetAccessToken(appid string, secret string) {
	var model models.WechatAccess
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret + "")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &model)
	if err != nil {
		panic(err)
	}
	fmt.Println(model.AccessToken)
	AccessToken = model.AccessToken
}

//GetAccessTokenTimer 获取微信AccessToken second 秒数
func GetAccessTokenTimer(appid string, secret string, second time.Duration) {
	timer1 := time.NewTicker(second * time.Second)
	for {
		select {
		case <-timer1.C:
			GetAccessToken(appid, secret)
		}
	}
}
