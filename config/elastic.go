package config

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

var Cfg elasticsearch.Config = elasticsearch.Config{
	Addresses: []string{
		"http://elastic:9200",
	},
	Transport: &http.Transport{
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second,
		DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS11,
		},
	},
}
