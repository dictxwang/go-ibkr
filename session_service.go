package ibkr

import (
	"encoding/json"
)

type SessionServiceI interface {
	PostAuthStatus() (*IServerItem, error)
	PostPingServer() (*PingServerResponse, error)
}

// SessionService :
type SessionService struct {
	client *Client
}

type PingServerResponse struct {
	Session    string            `json:"session"`
	SsoExpires int               `json:"ssoExpires"`
	Collission bool              `json:"collission"`
	UserId     int               `json:"userId"`
	Hmds       map[string]string `json:"hmds,omitempty"`
}

type ServerInfo struct {
	ServerName    string `json:"serverName"`
	ServerVersion string `json:"serverVersion"`
}

type IServerItem struct {
	Authenticated bool       `json:"authenticated"`
	Competing     bool       `json:"competing"`
	Connected     bool       `json:"connected"`
	Message       string     `json:"message"`
	MAC           string     `json:"MAC"`
	ServerInfo    ServerInfo `json:"serverInfo"`
	HardwareInfo  string     `json:"hardware_info"`
	Fail          string     `json:"fail"`
}

func (s *SessionService) PostAuthStatus() (*IServerItem, error) {

	var (
		res IServerItem
	)

	param := map[string]string{}
	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postJSON("/iserver/auth/status", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *SessionService) PostPingServer() (*PingServerResponse, error) {

	var (
		res PingServerResponse
	)

	param := map[string]string{}
	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postJSON("/tickle", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
