package models

type TokenResponse struct {
	AccessToken string `example:"eyJhbGciOiJIUzI1N74VlQTsu4TStQiloLfN0ENf6d64Ar0uSbrvtbUqFMGQ" json:"access_token"`
	ExpiresIn   int    `example:"18000" json:"expires_in"`
	TokenType   string `example:"Bearer" json:"token_type"`
} // @name token_response
