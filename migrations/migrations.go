package migrations

import (
	"github.com/nagahshi/bankApi/helpers"
	"github.com/nagahshi/bankApi/interfaces"
)

func createAccounts()  {
	db := helpers.ConnectDB()
	defer db.Close()

	users := &[2]interfaces.User{
		{Username: "Mikel", Email: "mikel@dark.com"},
		{Username: "Jonah", Email: "jonah@dark.com"},
	}

	for i := 0; i < len(users); i ++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))

		user := &interfaces.User{
			Username: users[i].Username,
			Email: users[i].Email,
			Password: generatedPassword,
		}

		db.Create(&user)

		account := interfaces.Account{
			Type: "Daily Account",
			Name: string(users[i].Username + "'s account"),
			Balance: uint(10000 * int(i * 1)),
			UserID: user.ID,
		}

		db.Create(&account)
	}
}

func Migrate()  {
	db := helpers.ConnectDB()
	defer db.Close()

	db.AutoMigrate(
		&interfaces.User{},
		&interfaces.Account{},
		)

	createAccounts()
}