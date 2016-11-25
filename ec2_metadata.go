package ec2metadata

import (
	"io/ioutil"
	"net/http"
)

const (
	addr = "http://169.254.169.254/latest/meta-data/"
)

// InstanceID returns the instance ID of the running instance.
func InstanceID() (string, error) {
	return get("instance-id")
}

// LocalIPAddress returns the local IP address of the running instance.
func LocalIPAddress() (string, error) {
	return get("local-ipv4")
}

// PublicIPAddress returns the public IP address of the running instance.
func PublicIPAddress() (string, error) {
	return get("public-ipv4")
}

func get(part string) (string, error) {
	resp, err := http.Get(addr + part)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
