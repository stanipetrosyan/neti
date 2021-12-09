package domain

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	State       string `json:"state"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}
