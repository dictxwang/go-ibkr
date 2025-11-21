package ibkr

import (
	"encoding/json"
)

type AuthServiceI interface {
	PostAuthStatus() (*AuthStatusResponse, error)
}

// AuthService :
type AuthService struct {
	client *Client
}

type AuthStatusServerInfo struct {
	ServerName    string `json:"serverName"`
	ServerVersion string `json:"serverVersion"`
}

type AuthStatusResponse struct {
	Authenticated bool                 `json:"authenticated"`
	Competing     bool                 `json:"competing"`
	Connected     bool                 `json:"connected"`
	Message       string               `json:"message"`
	MAC           string               `json:"MAC"`
	ServerInfo    AuthStatusServerInfo `json:"serverInfo"`
	HardwareInfo  string               `json:"hardware_info"`
	Fail          string               `json:"fail"`
}

func (s *AuthService) PostAuthStatus() (*AuthStatusResponse, error) {

	var (
		res AuthStatusResponse
	)

	param := map[string]string{}
	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postJSON("/v1/api/iserver/auth/status", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
