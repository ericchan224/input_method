package common

import (
	"net/http"
	"time"
)

// HTTPCli 定义一个全局的HTTP客户端
var HTTPCli = &http.Client{
	Timeout: time.Duration(15) * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:          10,
		MaxIdleConnsPerHost:   10,
		MaxConnsPerHost:       10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}
