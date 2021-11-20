package fetch

import (
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	bytes, err := http.Get(url)
	defer bytes.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(bytes.Body)
}
