package real

import (
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	Contents string
	Timeout  time.Duration
}

func (r *Retriever) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		return ""
	}
	return string(dumpResponse)

}
