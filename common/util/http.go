package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// PostJSON
//
//	@param url
//	@param data
//	@param header
//	@return res
//	@return err
func PostJSON(url string, data interface{}, header map[string]string) (res []byte, err error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return res, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(buf))
	if err != nil {
		return res, err
	}
	request.Header.Set("Content-Type", "application/json")
	for key, value := range header {
		request.Header.Set(key, value)
	}
	client := &http.Client{}
	client.Timeout = time.Second * 5
	resp, err := client.Do(request)
	if err != nil {
		return res, err
	}
	defer func() { _ = resp.Body.Close() }()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return respData, nil
}
