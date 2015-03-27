package http_consumer

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
)

func MakeHttpRequest(tcUrl string, queryUri string, authToken string, tlsConfig *tls.Config, proxy func(*http.Request) (*url.URL, error)) (*http.Response, error) {
	trafficControllerUrl, err := url.ParseRequestURI(tcUrl)

	if err != nil {
		return nil, err
	}

	host, _, err := net.SplitHostPort(trafficControllerUrl.Host)

	if err != nil {
		return nil, err

	}

	scheme := "https"

	if trafficControllerUrl.Scheme == "ws" {
		scheme = "http"
	}

	recentPath := fmt.Sprintf("%s://%s/%s", scheme, host, queryUri)

	transport := &http.Transport{Proxy: proxy, TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest("GET", recentPath, nil)
	req.Header.Set("Authorization", authToken)

	return client.Do(req)
}
