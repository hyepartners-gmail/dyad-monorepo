package auth

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

// GenerateJWT creates a JWT token for authenticated users.
func GenerateJWT(userID string, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})
	return token.SignedString(secretKey)
}

// OAuthConfig returns an OAuth2 config struct.
func OAuthConfig(clientID, clientSecret, redirectURL string, scopes []string, endpoint oauth2.Endpoint) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     endpoint,
	}
}
