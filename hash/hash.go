package hash

import (
	"github.com/Nikhil12894/gqlwithgo/graph/model"
	"golang.org/x/crypto/bcrypt"
)

//HashPassword hashes given password
func HashPassword(password string) string {
	if len(password) == 0 {
		return password

	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Authenticate(user model.User, hashedPassword string) bool {
	return CheckPasswordHash(user.Password, hashedPassword)
}
func AuthenticateLogin(user *model.Login, hashedPassword string) bool {
	return CheckPasswordHash(user.Password, hashedPassword)
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}
