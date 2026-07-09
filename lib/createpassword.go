package lib

import (
	"authenticatiion-flow/utils"
	"fmt"
)

func CreatePassword() string {
	var password1 string
	var password2 string

	fmt.Printf("Enter a strong password :  ")
	fmt.Scanf("%s", &password1)

	fmt.Printf("Confirm your password:  ")
	fmt.Scanf("%s", &password2)

	if password1 != password2 {
		utils.Clear()
		fmt.Printf("Wrong confirm password, press enter to back!  \n\n")
		fmt.Scanf("\n")
		CreatePassword()
	}

	return password1
}
