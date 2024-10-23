package helpers

import (
	_ "context"
	_ "fmt"
	_ "log"
	_ "os"
	_ "time"
	"golang/database"
	_ "github.com/dgrijalva/jwt-go"
)

func GenerateToken(user *database.User) (string, error) {

	
}	
	func ValidateToken(tokenString string) (*database.User, error) {
		return nil, nil
	}
