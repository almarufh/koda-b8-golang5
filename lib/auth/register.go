package auth

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/lib/utils"
	"fmt"
)

func Register(users *[]database.Users) (database.Users, string) {
	type results struct {
		message string
		code    string
		data    database.Users
	}
	var first string
	var last string
	var email string

	utils.Clear()
	fmt.Printf("\n\n--- Register ---\n\nWhat is your first name :  ")
	fmt.Scanf("%s", &first)

	fmt.Printf("What is your last name :  ")
	fmt.Scanf("%s", &last)

	fmt.Printf("What is your email :  ")
	fmt.Scanf("%s", &email)

	password := CreatePassword()

	newUser := database.Users{
		First:    first,
		Last:     last,
		Email:    email,
		Password: utils.Encrypted(password),
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nEmail belum terdaftar : %v\n", r)
			fmt.Printf("\nPress enter to back ...  ")
			fmt.Scanf("\n")
			main()
		}
	}()
	utils.Clear()
	var isAccept string
	fmt.Printf("\n\nIs it true?\n\n")
	fmt.Println("First Name : ", first)
	fmt.Println("Last Name : ", last)
	fmt.Println("Email : ", email)
	fmt.Printf("\n\nContinue (y/n) :  ")
	fmt.Scanf("%s", &isAccept)
	if isAccept == "y" {
		*users = append(*users, newUser)
		fmt.Printf("\n\nRegister success, press enter to back..")
		fmt.Scanf("\n")
	} else if isAccept == "n" {
		Register(users)
	} else {
		panic("Input wrong, press enter to back home..")
	}
	return newUser
}
