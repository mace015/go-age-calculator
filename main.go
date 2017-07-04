package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func fetchUserInput(question string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question, " ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error, please try again!")
	}
	inputTrim := strings.Trim(input, "\n")
	inputToInt, err := strconv.Atoi(inputTrim)
	return inputToInt
}

func main() {
	fmt.Println("Welcome to my first Go application!")

	day := fetchUserInput("On which day of the month were you born:")
	month := fetchUserInput("In what month where you born:")
	year := fetchUserInput("In what year where you born:")

	birthday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	timeSinceBirth := time.Now().Sub(birthday)
	yearsSinceBirthday := math.Floor(timeSinceBirth.Hours() / 8760)
	fmt.Println("According to my calculations, you are ", yearsSinceBirthday, " years old!")
	return
}
