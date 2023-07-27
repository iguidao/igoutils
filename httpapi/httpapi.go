package httpapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func PostJson(url string, jsonStr []byte, HeaderData map[string]string, timeout time.Duration) (bool, string) {

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	for headerkey, headervalue := range HeaderData {
		// req.Header.Set("Content-Type", "application/json")
		req.Header.Set(headerkey, headervalue)
	}
	client := &http.Client{Timeout: timeout * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil {
		errinfo := "post fails: " + err.Error()
		return false, errinfo
	}
	if resp.StatusCode != 200 {
		errinfo := "post fails, status: " + resp.Status
		return false, errinfo
	}
	defer resp.Body.Close()

	// statuscode := resp.StatusCode
	// hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	return true, string(body)

}

func GetDefault(url string, UriData map[string]string, HeaderData map[string]string, timeout time.Duration) (bool, string) {
	req, _ := http.NewRequest("GET", url, nil)
	for headerkey, headervalue := range HeaderData {
		// req.Header.Set("Content-Type", "application/json")
		req.Header.Set(headerkey, headervalue)
	}
	if UriData != nil {
		q := req.URL.Query()
		for urikey, urivalue := range UriData {
			q.Add(urikey, urivalue)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{Timeout: timeout * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil {
		errinfo := "get fails: " + err.Error()
		return false, errinfo
	}
	if resp.StatusCode != 200 {
		errinfo := "get fails, status: " + resp.Status
		return false, errinfo
	}
	defer resp.Body.Close()
	// statuscode := resp.StatusCode
	// hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	return true, string(body)
}
func PutDefault(url string, UriData map[string]string, HeaderData map[string]string, timeout time.Duration) (bool, string) {
	req, _ := http.NewRequest("PUT", url, nil)
	for headerkey, headervalue := range HeaderData {
		// req.Header.Set("Content-Type", "application/json")
		req.Header.Set(headerkey, headervalue)
	}
	if UriData != nil {
		q := req.URL.Query()
		for urikey, urivalue := range UriData {
			q.Add(urikey, urivalue)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{Timeout: timeout * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil {
		errinfo := "put fails: " + err.Error()
		return false, errinfo
	}
	if resp.StatusCode != 200 {
		errinfo := "put fails, status: " + resp.Status
		return false, errinfo
	}
	defer resp.Body.Close()
	// statuscode := resp.StatusCode
	// hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	return true, string(body)
}

func DeleteDefault(url string, UriData map[string]string, HeaderData map[string]string, timeout time.Duration) (bool, string) {
	req, _ := http.NewRequest("DELETE", url, nil)
	for headerkey, headervalue := range HeaderData {
		// req.Header.Set("Content-Type", "application/json")
		req.Header.Set(headerkey, headervalue)
	}
	if UriData != nil {
		q := req.URL.Query()
		for urikey, urivalue := range UriData {
			q.Add(urikey, urivalue)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{Timeout: timeout * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil {
		errinfo := "delete fails: " + err.Error()
		return false, errinfo
	}
	if resp.StatusCode != 200 {
		errinfo := "delete fails, status: " + resp.Status
		return false, errinfo
	}
	defer resp.Body.Close()
	// statuscode := resp.StatusCode
	// hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	return true, string(body)
}
