package main

// Resources at page 128

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "writeAGoodOne"
	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Fatalf("Password could not be hashed")
	}

	fmt.Printf("%s --> %s", password, hashedPassword)
	fmt.Println()

	err = decryptPassword(hashedPassword, password)
	if err != nil {
		log.Fatalf("Password does not match.")
	} else {
		fmt.Println("Decrypt password:")
		fmt.Printf("%s --> %s", hashedPassword, password)
	}

}

func hashPassword(s string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func decryptPassword(hs string, s string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hs), []byte(s))
	return err
}
