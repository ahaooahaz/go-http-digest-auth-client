package digest_auth_client

import (
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	username = "ahaoo"
	password = "test123"
	method   = "GET"
	uri      = "http://172.16.1.5"
)

func TestManual(t *testing.T) {
	var resp *http.Response
	var body []byte
	var err error

	dr := NewRequest(username, password, method, uri, "")

	if resp, err = dr.Execute(); err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		t.Fatal(err.Error())
	}

	t.Log(string(body))
}
