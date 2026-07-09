package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
	"os"
)

func Dashboard() {
	utils.Clear()
	getActived := database.GetActived()
	actived := *getActived
	var name string
	var indexActived int
	for i := range actived {
		if actived[i].Status == true {
			name = actived[i].Name
			indexActived = i
		}
	}
	var input int
	fmt.Printf("\n--- Welcome to system ---\n\nHello %s\n\n1. List All Users\n2. Logout\n\n0. Exit\n\nChoose a menu :   ", name)
	fmt.Scanf("%d", &input)
	if input == 1 {
		ListUsers()
	}
	if input == 2 {
		actived[indexActived].Status = false
		Menu()
	}

	if input == 0 {
		os.Exit(1)
	}
}
