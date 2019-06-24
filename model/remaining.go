package model

import (
	"net/http"
	"strings"
)

// Remaining : 요청 수 제한 정보
type Remaining struct {
	Group string
	Min   string
	Sec   string
}

func RemainingFromHeader(header http.Header) *Remaining {
	r := &Remaining{}

	v1s := strings.Split(header.Get("Remaining-Req"), ";")
	for _, v1 := range v1s {
		v2s := strings.Split(strings.TrimSpace(v1), "=")
		switch v2s[0] {
		case "group":
			r.Group = v2s[1]
		case "min":
			r.Min = v2s[1]
		case "sec":
			r.Sec = v2s[1]
		}
	}

	return r
}
