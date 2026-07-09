package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func ChangePassword() {
	utils.Clear()
	var email string
	fmt.Printf("\n--- [Forgot Password ] ---\n\nInput your email :  ")
	fmt.Scanf("%s", &email)
	getUsers := database.GetUsers()
	users := *getUsers

	var indexUser int
	var found bool = false

	for i := range users {
		if users[i].GetEmail() == email {
			found = true
			indexUser = i
			break
		}
	}

	if !found {
		fmt.Printf("\nEmail tidak ditemukan, tekan enter untuk kembali...")
		fmt.Scanf("\n")
		Menu()
		return
	}

	utils.Clear()
	fmt.Printf("--- CHANGE PASSWORD FOR %s ---\n\n", email)

	newPassword := CreatePassword()

	users[indexUser].Password = utils.Encrypted(newPassword)
	*getUsers = users

	utils.Clear()
	fmt.Printf("Password changed successfully! Press enter to home menu...")
	fmt.Scanf("\n")
	Menu()
}
