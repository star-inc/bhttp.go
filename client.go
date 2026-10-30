// bhttp.go
// (c) 2023 Star Inc.

package bHttp

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	DefaultUserAgent = "Mozilla/5.0 (compatible; Star Butterfly HTTP/1.0; +https://github.com/star-inc/bhttp.go)"
)

type Client struct {
	baseURL      *url.URL
	userAgent    string
	appendHeader http.Header
}

func NewHttpClient(baseURL string) (*Client, error) {
	httpClient := new(Client)
	httpClient.Initialize()
	if err := httpClient.SetBaseURL(baseURL); err != nil {
		return nil, err
	}
	return httpClient, nil
}

func (c *Client) Initialize() *Client {
	c.appendHeader = http.Header{}
	return c
}

func (c *Client) AddHeader(name string, values []string) *Client {
	c.appendHeader[name] = values
	return c
}

func (c *Client) SetBaseURL(baseURL string) error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	c.baseURL = u
	return nil
}

func (c *Client) SetUserAgent(userAgent string) *Client {
	c.userAgent = userAgent
	return c
}

func (c *Client) baseURLGlue(uri string) (string, error) {
	urlValues, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	if urlValues.Scheme != "" {
		return uri, nil
	}
	if c.baseURL == nil {
		return "", errors.New("baseURL is not set")
	}
	newURL := *c.baseURL
	newURL.Path = path.Join(newURL.Path, urlValues.Path)
	newURL.RawQuery = urlValues.RawQuery
	newURL.RawFragment = urlValues.RawFragment
	return newURL.String(), nil
}

func (c *Client) initRequest(method, fullURI string, data io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, fullURI, data)
	if err != nil {
		return nil, err
	}
	if data != nil {
		if _, ok := data.(*bytes.Buffer); ok {
			request.Header.Add("Content-Type", "application/json; charset=utf-8")
		} else if _, ok := data.(*strings.Reader); ok {
			request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		}
	}
	if c.userAgent != "" {
		request.Header.Add("User-Agent", c.userAgent)
	} else {
		request.Header.Add("User-Agent", DefaultUserAgent)
	}
	for key, value := range c.appendHeader {
		request.Header[key] = value
	}
	return request, nil
}

func (c *Client) Do(method, uri string, data io.Reader) (code StatusCode, body []byte, err error) {
	client := &http.Client{}
	var fullURI string
	fullURI, err = c.baseURLGlue(uri)
	if err != nil {
		return 0, nil, err
	}
	var request *http.Request
	request, err = c.initRequest(method, fullURI, data)
	if err != nil {
		return 0, nil, err
	}
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		if closeErr := response.Body.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return StatusCode(response.StatusCode), nil, err
	}
	return StatusCode(response.StatusCode), body, nil
}

func (c *Client) GET(uri string) (StatusCode, []byte, error) {
	return c.Do("GET", uri, nil)
}

func (c *Client) POST(uri string, data io.Reader) (StatusCode, []byte, error) {
	return c.Do("POST", uri, data)
}

func (c *Client) PUT(uri string, data io.Reader) (StatusCode, []byte, error) {
	return c.Do("PUT", uri, data)
}

func (c *Client) DELETE(uri string, data io.Reader) (StatusCode, []byte, error) {
	return c.Do("DELETE", uri, data)
}

func (c *Client) PATCH(uri string, data io.Reader) (StatusCode, []byte, error) {
	return c.Do("PATCH", uri, data)
}
