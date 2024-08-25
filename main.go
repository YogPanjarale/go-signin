package main

import (
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/tealeg/xlsx"
)

// Custom error types
var (
	ErrUsernameNotFound = fmt.Errorf("username not found")
	ErrPasswordIncorrect = fmt.Errorf("password incorrect")
)

// Check credentials against the Excel database
func checkCredentials(username, password string) error {
	// Opening the Excel file
	f, err := xlsx.OpenFile("users.xlsx")
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}

	// Loop through the rows in the Excel sheet
	for _, sheet := range f.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) < 2 {
				continue
			}

			if row.Cells[0].String() == username {
				// Compare stored hashed password with the provided password
				err := bcrypt.CompareHashAndPassword([]byte(row.Cells[1].String()), []byte(password))
				if err == nil {
					return nil // Credentials are valid
				}
				return ErrPasswordIncorrect
			}
		}
	}

	return ErrUsernameNotFound
}

// Sign-in handler function
func signInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		err := checkCredentials(username, password)
		if err == nil {
			fmt.Fprintf(w, "Welcome, %s!", username)
		} else if err == ErrUsernameNotFound {
			http.Error(w, "Username not found", http.StatusUnauthorized)
		} else if err == ErrPasswordIncorrect {
			http.Error(w, "Incorrect password", http.StatusUnauthorized)
		} else {
			http.Error(w, "An error occurred", http.StatusInternalServerError)
		}
	} else {
		http.ServeFile(w, r, "signin.html")
	}
}

// Main function to set up the HTTP server
func main() {
	http.HandleFunc("/signin", signInHandler)
	fmt.Println("Server is listening on port 8080...")
	fmt.Println("Visit http://localhost:8080/signin")
	http.ListenAndServe(":8080", nil)
}
