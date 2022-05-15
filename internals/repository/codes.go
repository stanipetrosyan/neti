package repository

type AuthorizationCode struct {
	ClientId string
	Code     string
}

type Codes interface {
	Add(code AuthorizationCode) bool
	FindBy(clientId string) string
}

type PostgresCodes struct{}

func (c *PostgresCodes) Add(code AuthorizationCode) bool {
	return true
}

func (c *PostgresCodes) FindBy(clientId string) string {
	return "ciao"
}
