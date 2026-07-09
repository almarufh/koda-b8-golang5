package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func ListUsers() {
	utils.Clear()
	users := database.GetUsers()
	user := *users
	for x := range user {
		urut := x + 1
		name := user[x].FullName()
		email := user[x].GetEmail()
		fmt.Printf("\n%d. %s [email : %s]", urut, name, email)
	}
	fmt.Printf("\n\npress enter to back Dashboard...  ")
	fmt.Scanf("\n")
	Dashboard()
}
