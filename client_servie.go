package ibkr

// ClientServiceI :
type ClientServiceI interface {
	Auth() AuthServiceI
}

// ClientService :
type ClientService struct {
	client *Client
}

// Auth :
func (s *ClientService) Auth() AuthServiceI {
	return &AuthService{s.client}
}

func (c *Client) Service() ClientServiceI {
	return &ClientService{c}
}
