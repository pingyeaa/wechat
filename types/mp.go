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
