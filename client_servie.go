package ibkr

// ClientServiceI :
type ClientServiceI interface {
	Session() SessionServiceI
	Account() AccountServiceI
	Contract() ContractServiceI
}

// ClientService :
type ClientService struct {
	client *Client
}

// Session :
func (s *ClientService) Session() SessionServiceI {
	return &SessionService{s.client}
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
