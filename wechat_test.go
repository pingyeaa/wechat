package wechat

import (
	"log"
	"testing"
)

func TestMP(t *testing.T) {
	c := NewMPConfig("123", "123")
	mp := NewMP(c)
	resp, err := mp.Auth().Code2Session("123")
	if err != nil {
		t.Error(err)
	}
	log.Println(resp.ErrorMsg)
}
