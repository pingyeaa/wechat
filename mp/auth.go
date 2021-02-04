package mp

import (
	"fmt"
	"wechat/types"
	"wechat/util/request"
)

// Auth .
type Auth struct {
	conf *Config
}

// Code2Session .
func (t *Auth) Code2Session(code string) (resp *types.MPCode2SessionResp, err error) {
	url := GetRequestHost() + fmt.Sprintf("/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", t.conf.AppID, t.conf.Secret, code)
	err = request.Get(url, &resp)
	if err != nil {
		return nil, fmt.Errorf("http request error: %s", err.Error())
	}
	return resp, err
}
