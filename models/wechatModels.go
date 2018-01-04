package models

//WechatAccess f
type WechatAccess struct {
	AccessToken string `json:"access_token"`
	WxpiresIn   int    `json:"expires_in"`
	Openid      string `json:"openid"`
	Scope       string `json:"scope"`
}
