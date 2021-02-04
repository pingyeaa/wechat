package mp

// MP .
type MP struct {
	C *Config
}

// Auth .
func (t *MP) Auth() *Auth {
	return &Auth{t.C}
}

// GetRequestHost .
func GetRequestHost() string {
	return "https://api.weixin.qq.com"
}
