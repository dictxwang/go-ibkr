package ibkr

// ClientServiceI :
type ClientServiceI interface {
	Auth() AuthServiceI
	Account() AccountServiceI
	Contract() ContractServiceI
}

// ClientService :
type ClientService struct {
	client *Client
}

// Auth :
func (s *ClientService) Auth() AuthServiceI {
	return &AuthService{s.client}
}

func (s *ClientService) Account() AccountServiceI {
	return &AccountService{s.client}
}

func (s *ClientService) Contract() ContractServiceI {
	return &ContractService{s.client}
}

func (c *Client) Service() ClientServiceI {
	return &ClientService{c}
}
