package authcontroller

import "golang.org/x/crypto/bcrypt"

// compareHashPassword The function that checks whether a
// hashed password and user input are the same or not
func compareHashPassword(user, userInput []byte) error {
	err := bcrypt.CompareHashAndPassword(user, userInput)
	if err != nil {
		return err
	}
	return nil
}
