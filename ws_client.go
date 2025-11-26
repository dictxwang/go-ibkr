package ibkr

import "log"

const (
	DefaultWebsocketBaseURL        = "wss://localhost:5000"
	DefaultWebsocketPrefixEndpoint = "/v1/api/ws"
)

// WebSocketClient :
type WebSocketClient struct {
	debug          bool
	logger         *log.Logger
	baseURL        string
	prefixEndpoint string
	skipTLSVerify  bool
}

func (c *WebSocketClient) debugf(format string, v ...interface{}) {
	if c.debug {
		c.logger.Printf(format, v...)
	}
}

// NewWebsocketClient :
func NewWebsocketClient(wsBaseUrl string, wsPrefixEndpoint string, skipTlsVerify bool) *WebSocketClient {
	baseUrl := DefaultWebsocketBaseURL
	if wsBaseUrl != "" {
		baseUrl = wsBaseUrl
	}
	prefixEndpoint := DefaultWebsocketPrefixEndpoint
	if wsPrefixEndpoint != "" {
		prefixEndpoint = wsPrefixEndpoint
	}
	return &WebSocketClient{
		logger:         newDefaultLogger(),
		baseURL:        baseUrl,
		prefixEndpoint: prefixEndpoint,
		skipTLSVerify:  skipTlsVerify,
	}
}

// NewDefaultWebsocketClient :
func NewDefaultWebsocketClient() *WebSocketClient {
	return NewWebsocketClient("", "", true)
}

// WithDebug :
func (c *WebSocketClient) WithDebug(debug bool) *WebSocketClient {
	c.debug = debug
	return c
}

// WithLogger :
func (c *WebSocketClient) WithLogger(logger *log.Logger) *WebSocketClient {
	c.debug = true
	c.logger = logger
	return c
}

// WithBaseURL :
func (c *WebSocketClient) WithBaseURL(url string) *WebSocketClient {
	c.baseURL = url
	return c
}

// WithPrefixEndpoint :
func (c *WebSocketClient) WithPrefixEndpoint(prefixEndpoint string) *WebSocketClient {
	c.prefixEndpoint = prefixEndpoint
	return c
}

// WithSkipTLSVersify :
func (c *WebSocketClient) WithSkipTLSVersify(skipTlsVerify bool) *WebSocketClient {
	c.skipTLSVerify = skipTlsVerify
	return c
}

// ErrHandler :
type ErrHandler func(isWebsocketClosed bool, err error)
