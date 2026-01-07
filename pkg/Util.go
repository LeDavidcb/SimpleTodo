package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Ask() (Todo, error) {
	b := bufio.NewReader(os.Stdin)
	var response string
	var todo Todo
	todo.Date = time.Now()

	fmt.Println("Write down the task you want to save: ")
	response, err := b.ReadString('\n')
	if err != nil {
		return Todo{}, errors.New("Error at reading response")
	}
	todo.Desc = strings.TrimSpace(response)

	fmt.Println("Write down the due date: quantity:unit[d/m/y]")
	response, err = b.ReadString('\n')
	if err != nil {
		return Todo{}, errors.New("Error at reading response")
	}

	response = strings.TrimSpace(response)

	// Validate if input contains a colon
	if !strings.Contains(response, ":") {
		return Todo{}, errors.New("Invalid response format - missing ':'")
	}

	// Split input into quantity and unit
	r := strings.Split(response, ":")
	if len(r) != 2 {
		return Todo{}, errors.New("Invalid response format")
	}

	// Extract and parse quantity and unit
	quantityStr := strings.TrimSpace(r[0])
	unit := strings.TrimSpace(strings.ToLower(r[1]))

	// Convert quantity to an integer
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		return Todo{}, errors.New("Invalid quantity - must be a positive number")
	}

	// Adjust due date based on the unit
	switch unit {
	case "d":
		todo.Due = time.Now().AddDate(0, 0, quantity)
	case "m":
		todo.Due = time.Now().AddDate(0, quantity, 0)
	case "y":
		todo.Due = time.Now().AddDate(quantity, 0, 0)
	default:
		return Todo{}, errors.New("Invalid time unit")
	}

	return todo, nil
}
