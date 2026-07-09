package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func CheckUser(email string, password string) {
	getUsers := database.GetUsers()
	getActived := database.GetActived()
	users := *getUsers
	actived := *getActived

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v\n", r)
			fmt.Printf("\nPress enter to home menu ...  ")
			fmt.Scanf("\n")
			Menu()
			return
		}
	}()

	if len(users) < 1 {
		panic("Email tidak ditemukan")
	}

	for i := range actived {
		if actived[i].Email == email {
			actived[i].Status = true
			*getActived = actived

			utils.Clear()
			fmt.Printf("Login success, press enter to Dashboard..")
			fmt.Scanf("\n")
			Dashboard()
			return
		}
	}

	var loginSukses bool = false

	for x := range users {
		loginEmail := users[x].GetEmail()
		loginPassword := users[x].GetPassword()

		if loginEmail == email && loginPassword == password {
			status := database.UserActive{
				Name:     users[x].FullName(),
				Email:    loginEmail,
				Password: loginPassword,
				Status:   true,
			}
			actived = append(actived, status)
			*getActived = actived
			loginSukses = true
			break
		}
	}

	if loginSukses {
		utils.Clear()
		fmt.Printf("Login success, press enter to Dashboard..")
		fmt.Scanf("\n")
		Dashboard()
	} else {
		fmt.Printf("\n\nWrong email or password, press enter to home menu...")
		fmt.Scanf("\n")
		Menu()
	}
}
