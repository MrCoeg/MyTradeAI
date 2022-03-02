package pkg

import (
	"errors"
	"io"
	"net/http"
)

func CreateConnection(url string) io.ReadCloser {

	res, err := http.Get(url)
	LogIfError(err)
	if res.StatusCode != 200 {
		panic(errors.New("not connected"))
	}
	return res.Body
}
