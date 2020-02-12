package ziggeo

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/go-resty/resty/v2"
)

// Connect connect object
type Connect struct {
	application *Ziggeo
	baseURI     string
}

// NewConnect connect constructor
func NewConnect(application *Ziggeo, baseURI string) *Connect {
	connect := new(Connect)
	connect.application = application
	connect.baseURI = baseURI
	return connect
}

// Request request template method
func (c *Connect) Request(method, path string, data map[string]string, file string) ([]byte, error) {
	var (
		url      = c.baseURI + path
		authInfo = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.application.Token, c.application.PrivateKey)))
		client   = resty.New()
		response *resty.Response
		err      error
	)

	if method == "GET" || method == "DELETE" {
		request := client.R().
			SetQueryParams(data).
			SetHeader("Authorization", fmt.Sprintf("Basic %s", authInfo))

		if method == "GET" {
			response, err = request.Get(url)
		}
		if method == "DELETE" {
			response, err = request.Delete(url)
		}
	} else if method == "POST" {
		if file != "" {
			response, err = client.R().
				SetFile(c.randomVideoName(12), file). // random video name
				SetFormData(data).
				Post(url)
		} else {
			response, err = client.R().
				SetFormData(data).
				Post(url)
		}
	} else {
		err = fmt.Errorf("unsupported method")
	}

	return response.Body(), err
}

// Get get method
func (c *Connect) Get(path string, data map[string]string) ([]byte, error) {
	return c.Request("GET", path, data, "")
}

// GetJSON get method
func (c *Connect) GetJSON(path string, data map[string]string, result *interface{}) error {
	response, err := c.Get(path, data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response, result)
	return err
}

// Delete delete method
func (c *Connect) Delete(path string, data map[string]string) ([]byte, error) {
	return c.Request("DELETE", path, data, "")
}

// DeleteJSON delete method
func (c *Connect) DeleteJSON(path string, data map[string]string, result *interface{}) error {
	response, err := c.Delete(path, data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response, result)
	return err
}

// Post post method
func (c *Connect) Post(path string, data map[string]string, file string) ([]byte, error) {
	return c.Request("POST", path, data, file)
}

// PostJSON post json method
func (c *Connect) PostJSON(path string, data map[string]string, file string, result *interface{}) error {
	response, err := c.Post(path, data, file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response, result)
	return err
}

// randomVideoName generate random video name
func (c *Connect) randomVideoName(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b) + ".mp4"
}
