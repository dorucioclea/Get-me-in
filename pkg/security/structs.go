package security

type TokenResponse struct{
	AccessToken 	string 	`json:"access_token"`
	TokenType 		string	`json:"token_type"`
	ExpiresIn		int		`json:"expires_in"`
	RefreshToken 	string	`json:"refresh_token"`
}

type TokenClaims struct{
	Issuer 		string
	Subject 	string
	Audience 	string
	Expiration 	string
	NotBefore 	string
	IssuedAt 	string
	Id			string
}
