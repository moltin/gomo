package gomo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c Client) url(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", c.Endpoint, c.APIVersion, endpoint)
}

func (c Client) buildRequest(method string, endpoint string, body []byte) (*http.Request, error) {
	var err error

	req, err := http.NewRequest(method, c.url(endpoint), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer: %s", c.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", defaultUserAgent)

	return req, nil
}

func (c *Client) do(wrapper *wrapper) error {
	var err error

	var body []byte
	if wrapper.Body != nil {

		rb := response{
			Data: wrapper.Body,
		}
		rbj, err := json.Marshal(rb)
		if err != nil {
			return err
		}

		body = rbj
	}

	wrapper.Request, err = c.buildRequest(wrapper.Method, wrapper.Endpoint, body)
	if err != nil {
		return err
	}

	wrapper.ExecutionTime.Start()
	r, err := c.httpClient.Do(wrapper.Request)
	wrapper.ExecutionTime.End()

	if err != nil {
		return err
	}

	wrapper.StatusCode = r.StatusCode

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &wrapper.Response)
	c.Log(wrapper)

	if len(wrapper.Response.Errors) > 0 {
		return errors.New(wrapper.Response.Errors[0].Detail)
	}

	return nil
}

// Post makes a POST request to the API
func (c *Client) Post(endpoint string, resource interface{}) (wrapper, error) {
	wrapper := newWrapper("post", endpoint, resource)
	return wrapper, c.do(&wrapper)
}

// Get makes a GET request to the API
func (c *Client) Get(endpoint string, resource ...interface{}) (wrapper, error) {
	wrapper := newWrapper("get", endpoint, resource...)
	return wrapper, c.do(&wrapper)
}

// Delete makes a DELETE request to the API
func (c *Client) Delete(endpoint string) (wrapper, error) {
	wrapper := newWrapper("delete", endpoint)
	return wrapper, c.do(&wrapper)
}

// Put makes a PUT request to the API
func (c *Client) Put(endpoint string, resource interface{}) (wrapper, error) {
	wrapper := newWrapper("put", endpoint, resource)
	return wrapper, c.do(&wrapper)
}
