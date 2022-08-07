package config

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/themartes/erd/config/envparams"
)

var Cfg elasticsearch.Config = elasticsearch.Config{
	Addresses: []string{
		GetEnvValue(envparams.ElasticsearchURL),
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
