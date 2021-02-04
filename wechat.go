package wechat

import "wechat/mp"

// NewMPConfig .
func NewMPConfig(appID string, secret string) *mp.Config {
	return &mp.Config{
		AppID:  appID,
		Secret: secret,
	}
}

// NewMP .
func NewMP(c *mp.Config) *mp.MP {
	return &mp.MP{C: c}
}
