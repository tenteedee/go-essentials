package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInputFloat(prompt string) float64 {
	var input string
	var floatInput float64
	for {
		fmt.Print(prompt)
		fmt.Scan(&input)
		var err error
		floatInput, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("Invalid input: %s. Please enter a valid number.\n", input)
		} else if floatInput <= 0 {
			fmt.Printf("Invalid input: %.2f. Please enter a positive number.\n", floatInput)
		} else {
			break
		}
	}
	return floatInput
}

func GetInputString(prompt string) string {
	var input string
	for {
		fmt.Print(prompt)
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')

		if input != "" {
			break
		} else {
			fmt.Println("Input must not be empty!")
		}
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}

func GetInputStringError(prompt string) (string, error) {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)

	if input == "" {
		return "", errors.New("input must not be empty")
	}

	return input, nil
}

func ReadFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 0, errors.New("error reading from file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("error parsing value to float")
	}

	return value, nil
}

func WriteFloatToFile(filename string, balance float64) {
	os.WriteFile(filename, []byte(fmt.Sprintf("%.2f", balance)), 0644)
	fmt.Println("Writing balance to file...")
}
