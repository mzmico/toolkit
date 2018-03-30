package wxapp

import (
	"fmt"

	"github.com/mzmico/toolkit/wechat"
)

type Config struct {
	AppId  string
	Secret string
}

type Session struct {
	wechat.Error
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

var (
	jscode2session = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

func JavascriptCodeToSession(c *Config, code string) (session *Session, err error) {

	url := fmt.Sprintf(
		jscode2session,
		c.AppId,
		c.Secret,
		code,
	)

	session = new(Session)

	err = wechat.GET(url, session)

	return

}
