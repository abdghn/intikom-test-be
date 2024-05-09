package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"intikom-test-be/model"
	"time"
)

const JWT_SECRET = "test123"

// GenerateToken -> generates token
func GenerateToken(user *model.User) string {

	claims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 3).Unix(),
		"iat":  time.Now().Unix(),
		"data": user,
		"sub":  uint(user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(JWT_SECRET))
	return t

}

// ValidateToken --> validate the given token
func ValidateToken(token string) (*jwt.Token, error) {

	//2nd arg function return secret key after checking if the signing method is HMAC and returned key is used by 'Parse' to decode the token)
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})
}
