package wechat

import (
	"log"
	"testing"
)

func TestMP(t *testing.T) {
	c := NewMPConfig("wxf4b2db3e02348b87", "7fda64a108c2d7d9dbd144d056eb25b2")
	mp := NewMP(c)
	resp, err := mp.Auth().GetAccessToken()
	if err != nil {
		t.Error(err)
	}
	log.Println(resp.AccessToken)
}
