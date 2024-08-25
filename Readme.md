# Signin Page in Golang
I have made a simple webpage that prompts the user for a username and password on a signin page.

The webpage compares the hash of the entered username and password with the values stored in the `users.xlsx` file. If the username is found and the password is correct, the user is signed in. Otherwise, an error message is displayed.

Additionally, I have created a command line utility to manage users. This tool allows you to view existing users and add new users by specifying their username and hash.

Example Username : `bits` 
Example password `pilani`