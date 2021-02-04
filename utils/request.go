package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get .
func Get(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http request error: %s", err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body error: %s", err.Error())
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		return fmt.Errorf("json decode error: %s", err.Error())
	}
	return nil
}
