package lib

import (
	"authenticatiion-flow/utils"
	"fmt"
)

func Login() {
	utils.Clear()
	var inputEmail string
	var inputPassword string
	fmt.Printf("\n---[ Page Login ]---\n\nEnter your email: ")
	fmt.Scanf("%s", &inputEmail)
	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s", &inputPassword)
	fmt.Printf("\n--------------------\n")
	password := utils.Encrypted(inputPassword)
	CheckUser(inputEmail, password)
}
