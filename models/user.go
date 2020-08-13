package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type User struct {
	ID           string   `json:"id"`
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	Password         string   `json:"password"`
	ProfilePic   string   `json:"profilePic"`
	MembershipID string   `json:"membershipId"`
}

func (u *User) HashPassword(pass string) error {
	bytePassword := []byte(pass)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err !=nil{
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

func (u *User) GenToken() (*AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "youRJube",
	})

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil{
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		ExpiredAt: expiredAt,
	}, nil
}

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword:=[]byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
