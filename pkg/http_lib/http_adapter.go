package http_lib

import (
	"bytes"
	"fmt"
	"net/http"
)

func Post(url string, body []byte, headers map[string]string) (*http.Response, error){
	return HttpRequest(url, body, headers, "POST")
}

func Get(url string, body []byte, headers map[string]string) (*http.Response, error){
	return HttpRequest(url, body, headers, "GET")
}

func Put(url string, body []byte, headers map[string]string) (*http.Response, error){
	return HttpRequest(url, body, headers, "PUT")
}

func Patch(url string, body []byte, headers map[string]string) (*http.Response, error){
	return HttpRequest(url, body, headers, "PATCH")
}

func Delete(url string, body []byte, headers map[string]string) (*http.Response, error){
	return HttpRequest(url, body, headers, "DELETE")
}

func HttpRequest(url string, body []byte, headers map[string]string, method string) (*http.Response, error){
	post, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			fmt.Println("k:", k, "v:", v)
			post.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(post)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, &ErrorString{
			Reason: resp.Status,
			Code:   resp.StatusCode,
		}
	}
	return resp, nil
}