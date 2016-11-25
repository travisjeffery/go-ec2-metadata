package ec2metadata

import (
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"testing"
)

func init() {
	http.DefaultClient.Transport = newTransport()
}

const (
	expID              = "i-4c3dc25a"
	expLocalIPAddress  = "10.0.0.1"
	expPublicIPAddress = "1.3.3.7"
)

func TestInstanceID(t *testing.T) {
	id, err := InstanceID()
	if err != nil {
		t.Error(err)
	}
	if expID != id {
		t.Errorf("expected %s, got %s", expID, id)
	}
}

func TestLocalIPAddress(t *testing.T) {
	ip, err := LocalIPAddress()
	if err != nil {
		t.Error(err)
	}
	if expLocalIPAddress != ip {
		t.Errorf("expected %s, got %s", expLocalIPAddress, ip)
	}
}

func TestPublicIPAddress(t *testing.T) {
	ip, err := PublicIPAddress()
	if err != nil {
		t.Error(err)
	}
	if expPublicIPAddress != ip {
		t.Errorf("expected %s, got %s", expPublicIPAddress, ip)
	}
}

type transport struct {
}

func newTransport() http.RoundTripper {
	return &transport{}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	base := path.Base(req.URL.Path)
	var responseBody string
	switch base {
	case "instance-id":
		responseBody = expID
	case "local-ipv4":
		responseBody = expLocalIPAddress
	case "public-ipv4":
		responseBody = expPublicIPAddress
	}
	response.Body = ioutil.NopCloser(strings.NewReader(responseBody))
	return response, nil
}
