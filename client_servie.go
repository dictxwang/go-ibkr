package ibkr

// ClientServiceI :
type ClientServiceI interface {
	Session() SessionServiceI
	Account() AccountServiceI
	Contract() ContractServiceI
	Order() OrdersServiceI
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

func (s *ClientService) Order() OrdersServiceI {
	return &OrdersService{s.client}
}

func (c *Client) Service() ClientServiceI {
	return &ClientService{c}
}
