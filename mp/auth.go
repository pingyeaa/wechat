package mp

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"wechat/types"
	"wechat/utils"
)

// Auth .
type Auth struct {
	conf *Config
}

// Code2Session .
func (t *Auth) Code2Session(code string) (resp *types.MPCode2SessionResp, err error) {
	url := GetRequestHost() + fmt.Sprintf("/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", t.conf.AppID, t.conf.Secret, code)
	err = utils.Get(url, &resp)
	if err != nil {
		return nil, fmt.Errorf("http request error: %s", err.Error())
	}
	return resp, err
}

// GetAccessToken .
func (t *Auth) GetAccessToken() (resp *types.MPGetAccessTokenResp, err error) {
	url := GetRequestHost() + fmt.Sprintf("/cgi-bin/token?appid=%s&secret=%s&grant_type=client_credential", t.conf.AppID, t.conf.Secret)
	err = utils.Get(url, &resp)
	if err != nil {
		return nil, fmt.Errorf("http request error: %s", err.Error())
	}
	return resp, err
}

// GetPaidUnionID .
func (t *Auth) GetPaidUnionID(accessToken string, openid string) (resp *types.MPGetPaidUnionIDResp, err error) {
	url := GetRequestHost() + fmt.Sprintf("/wxa/getpaidunionid?access_token=%s&openid=%s", accessToken, openid)
	err = utils.Get(url, &resp)
	if err != nil {
		return nil, fmt.Errorf("http request error: %s", err.Error())
	}
	return resp, err
}

// DecryptUserInfo .
func (t *Auth) DecryptUserInfo(encryptedData string, sessionKey string, iv string) (resp *types.MPDecryptUserInfoResp, err error) {
	err = t.decrypt(encryptedData, sessionKey, iv, &resp)
	if err != nil {
		return nil, fmt.Errorf("decrypt encryptedData failed: %s", err.Error())
	}
	return resp, nil
}

// DecryptPhoneData .
func (t *Auth) DecryptPhoneData(encryptedData string, sessionKey string, iv string) (resp *types.MPDecryptPhoneDataResp, err error) {
	err = t.decrypt(encryptedData, sessionKey, iv, &resp)
	if err != nil {
		return nil, fmt.Errorf("decrypt encryptedData failed: %s", err.Error())
	}
	return resp, nil
}

// decrypt .
func (t *Auth) decrypt(encryptedData string, sessionKey string, iv string, v interface{}) (err error) {
	aesData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return fmt.Errorf("base64 decode encryptedData failed: %s", err.Error())
	}
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return fmt.Errorf("base64 decode sessionKey failed: %s", err.Error())
	}
	aesIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return fmt.Errorf("base64 decode iv failed: %s", err.Error())
	}
	plainText, err := utils.AESCBCDecrypt(aesData, aesKey, aesIV...)
	if err != nil {
		return fmt.Errorf("aes cbc descrypt failed: %s", err.Error())
	}
	err = json.Unmarshal(plainText, &v)
	if err != nil {
		return fmt.Errorf("json decode failed: %s", err.Error())
	}
	return nil
}
