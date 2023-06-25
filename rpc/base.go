package rpc

import "time"

const (
	defaultMaxIdleConnsPerHost   = 100
	defaultMaxIdleConns          = 100
	defaultDialTimeout           = 10 * time.Second
	defaultKeepAliveTimeout      = 30 * time.Second
	defaultIdleConnTimeout       = 90 * time.Second
	defaultTLSHandshakeTimeout   = 10 * time.Second
	defaultExpectContinueTimeout = 1 * time.Second
)
