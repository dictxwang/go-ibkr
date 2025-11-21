package ibkr

// ServiceI :
type ServiceI interface {
	Auth() AuthServiceI
}

// Service :
type Service struct {
	client *Client
}

// Market :
func (s *Service) Market() AuthServiceI {
	return &AuthService{s.client}
}
