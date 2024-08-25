package main

import (
	"fmt"
	"log"
	"github.com/tealeg/xlsx"
	"golang.org/x/crypto/bcrypt"
)

const (
	usersFile = "users.xlsx"
	sheetName = "Sheet1"
)

func main() {
	fmt.Println("Welcome to the user management tool!")

	for {
		fmt.Println("Please select an option:")
		fmt.Println("1. View users and their password hashes")
		fmt.Println("2. Add a new user")
		fmt.Println("3. Exit")
		var option int
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			log.Fatal(err)
		}

		switch option {
		case 1:
			viewUsers()
		case 2:
			addUser()
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}

		fmt.Println()
	}
}

func viewUsers() {
	file, err := xlsx.OpenFile(usersFile)
	if err != nil {
	    
		log.Fatal(err)
	}

	sheet, ok := file.Sheet[sheetName]
	if !ok {
		log.Fatalf("Sheet '%s' not found in the file", sheetName)
	}

	for _, row := range sheet.Rows {
		username := row.Cells[0].String()
		passwordHash := row.Cells[1].String()

		fmt.Printf("Username: %s, Password Hash: %s\n", username, passwordHash)
	}
}

func addUser() {
	file, err := xlsx.OpenFile(usersFile)
	if err != nil {
		log.Fatal(err)
	}

	sheet, ok := file.Sheet[sheetName]
	if !ok {
		log.Fatalf("Sheet '%s' not found in the file", sheetName)
	}

	var username, password string

	fmt.Print("Enter username: ")
	_, err = fmt.Scanln(&username)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter password: ")
	_, err = fmt.Scanln(&password)
	
	if err != nil {
		log.Fatal(err)
	}

	// Hashing the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	row := sheet.AddRow()
	row.AddCell().SetValue(username)
	row.AddCell().SetValue(string(hashedPassword))

	err = file.Save(usersFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User added successfully!")
}