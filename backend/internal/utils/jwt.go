package util

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var SecretKey []byte

func CreateToken(username string,role string)(string,error){

	err := godotenv.Load()
	if err != nil {
		return "",err
	}
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"username": username,
		"role": role,
		"exp": time.Now().Add(1 *time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(SecretKey)
    if err != nil {
        return "", err
    }
	return tokenString,nil

}

func VerifyToken(tokenString string) error {
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return SecretKey, nil
   })
  
   if err != nil {
      return err
   }
  
   if !token.Valid {
      return fmt.Errorf("invalid token")
   }
  
   return nil
}

func SetAuthCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     "access-token",
		Value:    token,
		HttpOnly: true,   
		Secure: false,               
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	
}

func ExtractTokenFromRequest(r *http.Request) string {
	cookie, err := r.Cookie("access-token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func ParseToken(tokenString string) (username string, role string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
		SecretKey = []byte(os.Getenv("SECRET_KEY"))
		return SecretKey, nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, _ = claims["username"].(string)
		role, _ = claims["role"].(string)
		return username, role, nil
	}

	return "", "", fmt.Errorf("invalid token")
}

