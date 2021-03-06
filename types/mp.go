package types

// Error .
type Error struct {
	ErrorCode int    `json:"errode"`
	ErrorMsg  string `json:"errmsg"`
}

// MPCode2Session .
type MPCode2SessionResp struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	Error
}

// MPGetAccessTokenResp .
type MPGetAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Error
}

// MPGetPaidUnionIDResp .
type MPGetPaidUnionIDResp struct {
	UnionID string `json:"unionid"`
	Error
}

// MPDecryptUserInfoResp .
type MPDecryptUserInfoResp struct {
	OpenID    string       `json:"openId"`
	NickName  string       `json:"nickName"`
	Gender    string       `json:"gender"`
	City      string       `json:"city"`
	Province  string       `json:"province"`
	Country   string       `json:"country"`
	AvatarURL string       `json:"avatarUrl"`
	UnionID   string       `json:"unionId"`
	Watermark *MPWaterMark `json:"watermark"`
}

// MPDecryptPhoneDataResp .
type MPDecryptPhoneDataResp struct {
	PhoneNumber     string       `json:"phoneNumber"`
	PurePhoneNumber string       `json:"purePhoneNumber"`
	CountryCode     string       `json:"countryCode"`
	WaterMark       *MPWaterMark `json:"watermark"`
}

// MPWaterMark 水印
type MPWaterMark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"`
}
