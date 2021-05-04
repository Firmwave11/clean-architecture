package utils

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//Component ....
type Component struct {
	URL    string
	Header map[string]string
	rTimeout
}

type rTimeout struct {
	duration time.Duration
	set      bool
}

//WithTimeout is setting....
func (r *Component) WithTimeout(d int64) Methods {
	r.rTimeout.duration = time.Duration(d) * time.Second
	r.rTimeout.set = true
	return r
}

//Methods is func to request
type Methods interface {
	Post(payload []byte) ([]byte, int, error)
	Get() ([]byte, int, error)
	Put(payload []byte) ([]byte, int, error)
	Delete(payload []byte) ([]byte, int, error)
}

func buf(p []byte) io.Reader {
	if p != nil {
		return bytes.NewBuffer(p)
	}
	return nil
}

func do(r *Component, payload []byte, method string) ([]byte, int, error) {

	req, err := http.NewRequest(method, r.URL, buf(payload))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	//Set header
	if r.Header != nil {
		for k, v := range r.Header {
			req.Header.Set(k, v)
		}
	}

	//Set timeout if exist
	if r.rTimeout.set {
		cons, cancel := context.WithTimeout(context.Background(), r.rTimeout.duration)
		req = req.WithContext(cons)
		defer cancel()
	}

	//Do request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	//Read result
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, res.StatusCode, nil
}

//New ...
func New(header map[string]string, url string) *Component {
	return &Component{
		URL:    url,
		Header: header,
	}
}

//Post is ...
func (r *Component) Post(payload []byte) ([]byte, int, error) {
	return do(r, payload, http.MethodPost)
}

//Get ...
func (r *Component) Get() ([]byte, int, error) {
	return do(r, nil, http.MethodGet)
}

//Delete ...
func (r *Component) Delete(payload []byte) ([]byte, int, error) {
	return do(r, payload, http.MethodDelete)
}

//Put ...
func (r *Component) Put(payload []byte) ([]byte, int, error) {
	return do(r, payload, http.MethodPut)
}
