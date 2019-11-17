package sendrequest

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"time"
)

// Post sends a POST request to a given URL with the given headers and body
func Post(headers map[string]string, url string, obj interface{}) (*http.Response, error) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(obj)

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	if _, ok := req.Header["Content-Type"]; !ok {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   20 * time.Second,
				KeepAlive: 20 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
