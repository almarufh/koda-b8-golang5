package main

import (
	"fmt"
	"os"
)

type userActive struct {
	name     string
	email    string
	password string
	status   bool
}

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

var user []Users = []Users{
	{
		first:    "Alma'ruf",
		last:     "Hidayat",
		email:    "almarufhidayat99@gmail.com",
		password: "1234",
	},
	{
		first:    "Ali",
		last:     "Ghufro",
		email:    "alhyghuron@gmail.com",
		password: "1234",
	},
	{
		first:    "CuanBot",
		last:     "",
		email:    "a",
		password: "a",
	},
}

var actived []userActive = []userActive{}

var status bool = false

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

func authLogin(email string, password string) {
	var activedSaved bool = false
	var indexActived int
	for i := range actived {
		if actived[i].email == email {
			activedSaved = true
			indexActived = i
		}
	}

	if activedSaved == true {
		actived[indexActived].status = true
		dashboard()
	}

	for x := range user {
		loginEmail := user[x].getEmail
		loginPassword := user[x].getPassword()
		if loginEmail() == email && loginPassword == password {
			status := userActive{
				name:     user[x].fullName(),
				email:    loginEmail(),
				password: loginPassword,
				status:   true,
			}
			actived = append(actived, status)
			dashboard()
		}
	}
	if status == true {
		clear()
		fmt.Printf("Login success, press enter to Dashboard..")
		fmt.Scanf("\n")
		dashboard()
	} else {
		fmt.Printf("\n\nWrong email or password, press enter to restart...")
		fmt.Scanf("\n")
		login()
	}
}

func dashboard() {
	clear()
	var name string
	var indexActived int
	for i := range actived {
		if actived[i].status == true {
			name = actived[i].name
			indexActived = i
		}
	}
	var input int
	fmt.Printf("\n--- Welcome to system ---\n\nHello %s\n\n1. List All Users\n2. Logout\n\n0. Exit\n\nChoose a menu :   ", name)
	fmt.Scanf("%d", &input)
	if input == 1 {
		listUsers()
	}
	if input == 2 {
		actived[indexActived].status = false
		main()
	}

	if input == 0 {
		exit()
	}
}

func listUsers() {
	clear()
	for x := range user {
		urut := x + 1
		name := user[x].fullName()
		email := user[x].getEmail()
		fmt.Printf("\n%d. %s %s", urut, name, email)
	}
	fmt.Printf("\n\npress enter to back Dashboard...  ")
	fmt.Scanf("\n")
	dashboard()
}

func login() {
	clear()
	var inputEmail string
	var inputPassword string
	fmt.Printf("\n\n---Login---\n\nEnter your email: ")
	fmt.Scanf("%s", &inputEmail)
	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s", &inputPassword)
	authLogin(inputEmail, inputPassword)
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
