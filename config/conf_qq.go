package config

import "fmt"

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` //登陆之后的回调地址
}

func (q QQ) GetPath() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("http://graph.qq.com/oauth2.0/show?which=Login&display=pc&reponse_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
