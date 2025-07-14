package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%^&*()-_"
	numberCharSet  = "1234567890"
	minSpecialChar = 4
	minUpperChar   = 4
	minNumberChar  = 4
	passwordLength = 30
)

func main() {

	totalCharLenWithoutLowerChar := minUpperChar + minSpecialChar + minNumberChar

	if totalCharLenWithoutLowerChar >= passwordLength {
		fmt.Println("Please provide valid password length")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("How many password you want to generate? - ")
	scanner.Scan()

	numberOfPassword, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Print("Please provide correct value for number of passwords \n")
		os.Exit(1)
	}

	rand.Seed(time.Now().Unix())

	for i := 0; i < numberOfPassword; i++ {
		password := generatePassword()
		fmt.Printf("password %v is %v \n", i+1, password)
	}

}

func generatePassword() string {

	password := ""

	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))

		password = password + string(specialCharSet[random])

	}

	for i := 0; i < minUpperChar; i++ {
		random := rand.Intn(len(upperCharSet))

		password = password + string(upperCharSet[random])
	}

	for i := 0; i < minNumberChar; i++ {
		random := rand.Intn(len(numberCharSet))

		password = password + string(numberCharSet[random])
	}

	totalCharLenWithoutLowerChar := minUpperChar + minSpecialChar + minNumberChar

	remainingCharLen := passwordLength - totalCharLenWithoutLowerChar

	for i := 0; i < remainingCharLen; i++ {
		random := rand.Intn(len(lowerCharSet))
		password = password + string(lowerCharSet[random])
	}

	passwordRune := []rune(password)
	rand.Shuffle(len(passwordRune), func(i, j int) {
		passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i]

	})
	password = string(passwordRune)
	return password
}
