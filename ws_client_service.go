package ibkr

import (
	"crypto/tls"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// WebsocketClientServiceI :
type WebsocketClientServiceI interface {
	Public(sessionToken string) (*WebsocketPublicService, error)
	PublicWithSourceIP(sessionToken string, sourceIP string) (*WebsocketPublicService, error)
	Private(sessionToken string) (*WebsocketPublicService, error)
	PrivateWithSourceIP(sessionToken string, sourceIP string) (*WebsocketPublicService, error)
}

// WebsocketClientService :
type WebsocketClientService struct {
	client *WebSocketClient
}

// Service :
func (c *WebSocketClient) Service() *WebsocketClientService {
	return &WebsocketClientService{c}
}

// Public :
func (s *WebsocketClientService) Public(sessionToken string) (*WebsocketPublicService, error) {
	url1 := s.client.baseURL + s.client.prefixEndpoint
	dialer := generateCustomDialer(s.client.skipTLSVerify, "")
	cookie := &http.Cookie{
		Name:   "api",
		Value:  sessionToken,
		Path:   "/",
		Domain: "localhost",
	}
	jar, _ := cookiejar.New(nil)
	serverURL, _ := url.Parse(s.client.baseURL)
	jar.SetCookies(serverURL, []*http.Cookie{cookie})
	dialer.Jar = jar
	requestHeader := makeRequestHeader(sessionToken)
	c, _, err := dialer.Dial(url1, requestHeader)
	if err != nil {
		return nil, err
	}

	//// TODO
	//login := fmt.Sprintf("{\"session\":\"%s\"}", sessionToken)
	//c.WriteMessage(websocket.TextMessage, []byte(login))
	return &WebsocketPublicService{
		client:     s.client,
		connection: c,
	}, nil
}

// PublicWithSourceIP :
func (s *WebsocketClientService) PublicWithSourceIP(sessionToken, sourceIP string) (*WebsocketPublicService, error) {
	url := s.client.baseURL + s.client.prefixEndpoint
	dialer := generateCustomDialer(s.client.skipTLSVerify, sourceIP)
	requestHeader := makeRequestHeader(sessionToken)
	c, _, err := dialer.Dial(url, requestHeader)
	if err != nil {
		return nil, err
	}
	return &WebsocketPublicService{
		client:     s.client,
		connection: c,
	}, nil
}

// Private :
func (s *WebsocketClientService) Private(sessionToken string) (*WebsocketPrivateService, error) {
	url := s.client.baseURL + s.client.prefixEndpoint
	dialer := generateCustomDialer(s.client.skipTLSVerify, "")
	requestHeader := makeRequestHeader(sessionToken)
	c, _, err := dialer.Dial(url, requestHeader)
	if err != nil {
		return nil, err
	}
	return &WebsocketPrivateService{
		client:     s.client,
		connection: c,
	}, nil
}

// PrivateWithSourceIP :
func (s *WebsocketClientService) PrivateWithSourceIP(sessionToken, sourceIP string) (*WebsocketPrivateService, error) {
	url := s.client.baseURL + s.client.prefixEndpoint
	dialer := generateCustomDialer(s.client.skipTLSVerify, sourceIP)
	requestHeader := makeRequestHeader(sessionToken)
	c, _, err := dialer.Dial(url, requestHeader)
	if err != nil {
		return nil, err
	}
	return &WebsocketPrivateService{
		client:     s.client,
		connection: c,
	}, nil
}

func generateCustomDialer(skipTlsVerify bool, sourceIP string) *websocket.Dialer {
	dialer := websocket.DefaultDialer
	tlsConfig := tls.Config{
		InsecureSkipVerify: skipTlsVerify,
	}
	dialer.TLSClientConfig = &tlsConfig
	if sourceIP != "" {
		dialer.NetDial = func(network, addr string) (net.Conn, error) {
			localAddr, err := net.ResolveTCPAddr(network, sourceIP+":0")
			if err != nil {
				return nil, err
			}

			remoteAddr, err := net.ResolveTCPAddr(network, addr)
			if err != nil {
				return nil, err
			}

			return net.DialTCP(network, localAddr, remoteAddr)
		}
	}
	return dialer
}

func makeRequestHeader(sessionToken string) http.Header {
	httpHeader := http.Header{}
	//httpHeader.Add("origin", "interactivebrokers.github.io")
	//httpHeader.Add("cookie", "api="+sessionToken)
	return httpHeader
}
