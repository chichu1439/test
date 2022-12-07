package models

import (
	"net/http"
	"strings"
)

type OutgressRequest struct {
	Body []byte `json:"body"`
}

type OutgressResponse struct {
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
}

func (o *OutgressResponse) GetHeaderIgnoreCase(key string) (value string) {

	for k, vs := range o.Headers {
		if len(vs) == 0 {
			return ""
		}
		if strings.EqualFold(k, key) {
			return vs[0]
		}
	}

	return ""
}
