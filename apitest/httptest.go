package main

import (
	"crypto/tls"
	"net/http"
	"strings"
)

var (
	//contentType
	C_JsonType    = 1
	C_XmlType     = 2
	C_TxtType     = 3
	C_JosnContent = "application/json"
	C_XmlContent  = "application/xml"
	C_TxtContent  = "application/txt"

	//methodType
	M_PostType      = 1
	M_GetType       = 2
	M_DelteType     = 3
	M_PutType       = 4
	M_PostContent   = "POST"
	M_GetContent    = "GET"
	M_PutContent    = "PUT"
	M_DeleteContent = "DELETE"
)

func GetClient() http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return http.Client{Transport: tr}
}

func SetRequester(contentType int, cookie string, url string, methodType int, data string) *http.Request {
	req, _ := http.NewRequest(SetMethodType(methodType), SetRequesterUrl(url), strings.NewReader(data))
	req.Header.Set("Content-Type", SetContentType(contentType))
	//req.Header.Set("Cookie", SetCookie(cookie))
	req.SetBasicAuth("token-9gqkf", "vfkltx6nbfwvvv8mbxtl5jn674q2xbcvnstzfvrgvwwcrv9s4btf59")
	//req.Header.Set("Authorization", "Basic dG9rZW4tOWdxa2Y6dmZrbHR4Nm5iZnd2dnY4bWJ4dGw1am42NzRxMnhiY3Zuc3R6ZnZyZ3Z3d2NydjlzNGJ0ZjU5")
	return req
}

func SetContentType(contentType int) string {
	switch contentType {
	case C_JsonType:
		return C_JosnContent
	case C_XmlType:
		return C_XmlContent
	case C_TxtType:
		return C_TxtContent
	}
	return C_JosnContent
}

func SetCookie(cookie string) string {
	return cookie
}

func SetMethodType(methodType int) string {
	switch methodType {
	case M_GetType:
		return M_GetContent
	case M_PostType:
		return M_PostContent
	case M_DelteType:
		return M_DeleteContent
	case M_PutType:
		return M_PutContent
	}
	return M_GetContent
}

func SetRequesterUrl(url string) string {
	return url
}
