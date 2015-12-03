package util

import (
	"net/http"
)

func BaseURL(r *http.Request) string {

	var protocol string

	if r.TLS != nil {
		protocol = "https"
	} else {
		protocol = "http"
	}

	host := r.Host

	return protocol + "://" + host
}
