package database

// Variable User
type Users struct {
	First    string
	Last     string
	Email    string
	Password string
}

// Variable User Actived
type UserActive struct {
	Name     string
	Email    string
	Password string
	Status   bool
}

// Type data results func
type Results struct {
	Message string
	Code    string
	Data    Users
}

var Project string = "s y s t e m"

// Func Output Memory Adress
var actived []UserActive = []UserActive{}

func GetActived() *[]UserActive {
	return &actived
}

var users []Users = []Users{
	{
		First:    "Alma'ruf",
		Last:     "Hidayat",
		Email:    "owner@gmail.com",
		Password: "81dc9bdb52d04dc20036dbd8313ed055",
	},
	{
		First:    "Alma'ruf",
		Last:     "Hidayat",
		Email:    "owner",
		Password: "81dc9bdb52d04dc20036dbd8313ed055",
	},
	{
		First:    "Ali",
		Last:     "Ghufro",
		Email:    "admin",
		Password: "81dc9bdb52d04dc20036dbd8313ed055",
	},
	{
		First:    "CuanBot",
		Last:     "",
		Email:    "a",
		Password: "0cc175b9c0f1b6a831c399e269772661",
	},
}

func GetUsers() *[]Users {
	return &users
}

// Func Get Email, Password, Full Name
func (u Users) FullName() string {
	return u.First + " " + u.Last
}

func (u Users) GetEmail() string {
	return u.Email
}

func (u Users) GetPassword() string {
	return u.Password
}
