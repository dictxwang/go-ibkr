package ibkr

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	DefaultBaseUrl        = "https://localhost:5000"
	DefaultPrefixEndpoint = "/v1/api"
)

// Client :
type Client struct {
	httpClient *http.Client

	debug  bool
	logger *log.Logger

	baseURL        string
	endpointPrefix string
	key            string
	secret         string

	referer string
}

func (c *Client) debugf(format string, v ...interface{}) {
	if c.debug {
		c.logger.Printf(format, v...)
	}
}

// NewClient :
func NewClient(restBaseUrl, restEndpointPrefix string, skipTlsVerify bool) *Client {

	baseUrl := DefaultBaseUrl
	if restBaseUrl != "" {
		baseUrl = restBaseUrl
	}

	endpointPrefix := DefaultPrefixEndpoint
	if restEndpointPrefix != "" {
		endpointPrefix = restEndpointPrefix
	}

	var httpClient *http.Client
	if skipTlsVerify {
		// Configure TLS client to skip verification
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true, // DO NOT USE IN PRODUCTION
		}
		// Create a custom HTTP client with the TLS configuration
		httpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}
	} else {
		httpClient = &http.Client{}
	}

	return &Client{
		httpClient:     httpClient,
		logger:         newDefaultLogger(),
		baseURL:        baseUrl,
		endpointPrefix: endpointPrefix,
	}
}

// NewDefaultClient :
func NewDefaultClient() *Client {
	return NewClient("", "", true)
}

// WithHTTPClient :
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient

	return c
}

// WithDebug :
func (c *Client) WithDebug(debug bool) *Client {
	c.debug = debug

	return c
}

// WithLogger :
func (c *Client) WithLogger(logger *log.Logger) *Client {
	c.debug = true
	c.logger = logger

	return c
}

// WithBaseURL :
func (c *Client) WithBaseURL(url string) *Client {
	c.baseURL = url

	return c
}

func (c *Client) WithReferer(referer string) *Client {
	c.referer = referer

	return c
}

// Request :
func (c *Client) Request(req *http.Request, dst interface{}) error {
	_, err := c.RequestFull(req, false, dst)
	return err
}

// RequestFull :
func (c *Client) RequestFull(req *http.Request, conciseResponse bool, dst interface{}) ([]byte, error) {
	c.debugf("request: %v", req)
	resp, err := c.httpClient.Do(req)
	c.debugf("response: %v", resp)
	if err != nil {
		return nil, err
	}
	c.debugf("response status code: %v", resp.StatusCode)
	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	switch {
	case 200 <= resp.StatusCode && resp.StatusCode <= 299:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		fmt.Printf("response body: %+v\n", string(body))
		c.debugf("response body: %v", string(body))

		if conciseResponse {
			return body, nil
		} else {
			if err := json.Unmarshal(body, &dst); err != nil {
				return body, err
			} else {
				return nil, nil
			}
		}
	case resp.StatusCode == http.StatusBadRequest:
		return nil, fmt.Errorf("%v: Need to send the request with GET / POST (must be capitalized) url=%s", ErrBadRequest, req.URL.String())
	case resp.StatusCode == http.StatusUnauthorized:
		return nil, fmt.Errorf("%w: invalid key/secret", ErrInvalidRequest)
	case resp.StatusCode == http.StatusForbidden:
		return nil, fmt.Errorf("%w: not permitted", ErrForbiddenRequest)
	case resp.StatusCode == http.StatusNotFound:
		return nil, fmt.Errorf("%w: wrong path", ErrPathNotFound)
	default:
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
}

func (c *Client) getPublic(path string, query url.Values, dst interface{}) error {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = c.endpointPrefix + path
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) deletePublic(path string, query url.Values, dst interface{}) error {

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = c.endpointPrefix + path
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return err
	}

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) postJSON(path string, body []byte, dst interface{}) error {

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = c.endpointPrefix + path

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) postJSONConciseResponse(path string, body []byte) ([]byte, error) {

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	u.Path = c.endpointPrefix + path

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	respBody, err := c.RequestFull(req, true, nil)
	return respBody, err
}

func (c *Client) postForm(path string, body url.Values, dst interface{}) error {

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil
	}
	u.Path = path

	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return err
	}

	if err := c.Request(req, &dst); err != nil {
		return err
	}

	return nil
}
