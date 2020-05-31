package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/nagahshi/bankApi/helpers"
	"github.com/nagahshi/bankApi/interfaces"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Login(username string, pass string) map[string]interface{} {
	db := helpers.ConnectDB()
	defer db.Close()

	user := &interfaces.User{}
	if db.Where("username = ?", username).Find(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if passErr != nil && passErr == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"message": "Senha inválida"}
	}

	var accounts []interfaces.ResponseAccount
	db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

	responseUser := &interfaces.ResponseUser{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Accounts: accounts,
	}

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry": time.Now().Add(time.Minute ^ 60).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)

	token, err := jwtToken.SignedString([]byte("TokenPasswordSecret"))
	helpers.HandleErr(err)

	return map[string]interface{}{
		"jwt": token,
		"data": responseUser,
	}
}
