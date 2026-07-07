package main

import (
	"fmt"
	"os"
)

type Login struct {
	email    string
	password string
}

type Users struct {
	first    string
	last     string
	email    string
	password string
}

var user = []Users{}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func (u Users) fullName() string {
	return u.first + " " + u.last
}

func (u Users) getEmail() string {
	return u.email
}

func (u Users) getPassword() string {
	return u.password
}

func register() {
	var first string
	var last string
	var email string
	var password string

	clear()
	fmt.Printf("\n\n--- Register ---\n\nWhat is your first name :  ")
	fmt.Scanf("%s", &first)

	fmt.Printf("What is your last name :  ")
	fmt.Scanf("%s", &last)

	fmt.Printf("What is your email :  ")
	fmt.Scanf("%s", &email)

	fmt.Printf("Enter a strong password :  ")
	fmt.Scanf("%s", &password)

	fmt.Printf("Confirm your password:  ")
	fmt.Scanf("%s", &password)

	new := Users{
		first:    first,
		last:     last,
		email:    email,
		password: password,
	}

	clear()
	var isAccept string
	fmt.Printf("\n\nIs it true?\n\n")
	fmt.Println("First Name : ", first)
	fmt.Println("Last Name : ", last)
	fmt.Println("Email : ", email)
	fmt.Printf("\n\nContinue (y/n) :  ")
	fmt.Scanf("%s", &isAccept)
	if isAccept == "y" {
		user = append(user, new)
		fmt.Printf("\n\nRegister success, press enter to back..")
		fmt.Scanf("\n")
		main()
	}

	if isAccept == "n" {
		register()
	}
	fullname := user[0].fullName()
	fmt.Println(fullname)
}

func exit() {
	os.Exit(0)
}

func login() {
	clear()
	var inputEmail string
	var inputPassword string
	fmt.Printf("\n\n---Login---\n\nEnter your email: ")
	fmt.Scanf("%s", &inputEmail)
	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s", &inputPassword)
	dataLogin := Login{
		email:    inputEmail,
		password: inputPassword,
	}
	fmt.Println(dataLogin)
}

func main() {
	var input int
	clear()
	fmt.Println("--- Welcome to system ---")
	fmt.Printf("\n1. Register\n2. Login\n3. Forgot Password\n\n0. Exit\n\n")
	fmt.Printf("Choose a menu : ")
	fmt.Scanf("%d", &input)
	if input == 1 {
		register()
	}

	if input == 2 {
		login()
	}

	if input == 0 {
		exit()
	}
}
