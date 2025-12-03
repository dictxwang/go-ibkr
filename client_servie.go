package ibkr

// ClientServiceI :
type ClientServiceI interface {
	Session() SessionServiceI
	Account() AccountServiceI
	Contract() ContractServiceI
	Order() OrdersServiceI
	OrderMonitoring() OrderMonitoringServiceI
	Portfolio() PortfolioServiceI
}

// ClientService :
type ClientService struct {
	client *Client
}

// Session :
func (s *ClientService) Session() SessionServiceI {
	return &SessionService{s.client}
}

// Account :
func (s *ClientService) Account() AccountServiceI {
	return &AccountService{s.client}
}

// Contract :
func (s *ClientService) Contract() ContractServiceI {
	return &ContractService{s.client}
}

// Order :
func (s *ClientService) Order() OrdersServiceI {
	return &OrdersService{s.client}
}

// OrderMonitoring :
func (s *ClientService) OrderMonitoring() OrderMonitoringServiceI {
	return &OrderMonitoringService{s.client}
}

// Portfolio :
func (s *ClientService) Portfolio() PortfolioServiceI {
	return &PortfolioService{s.client}
}

// Service :
func (c *Client) Service() ClientServiceI {
	return &ClientService{c}
}
