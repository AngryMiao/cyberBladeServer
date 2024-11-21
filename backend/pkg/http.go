package pkg

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func IsAllowHost(origin string, allowHosts []string) bool {
	for _, host := range allowHosts {
		if host == "*" || strings.HasSuffix(strings.ToLower(origin), host) {
			return true
		}
	}
	return false
}

func CloneURLValues(v url.Values) url.Values {
	v2 := make(url.Values, len(v))
	for k, vv := range v {
		v2[k] = append([]string(nil), vv...)
	}
	return v2
}

func HttpGet(url string) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "text/plain")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
