package httphelper

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HTTPHelper interface {
	PostForm(uri string, data url.Values) (resp *http.Response, err error)
	ResponseBodyAsBytes(resp *http.Response) ([]byte, error)
}

type ActualHTTPHelper struct {
}

func (a *ActualHTTPHelper) PostForm(uri string, data url.Values) (resp *http.Response, err error) {
	return http.PostForm(uri, data)

}

func (a *ActualHTTPHelper) ResponseBodyAsBytes(resp *http.Response) ([]byte, error) {
	if resp.Body == nil {
		return nil, errors.New("Nil response body")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
