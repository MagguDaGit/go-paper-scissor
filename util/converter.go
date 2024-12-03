package util

import (
	"errors"
	"fmt"
	"rubrumcreation.com/go-paper-scissor/models"
	"strconv"
	"strings"
)

func ConvertBool(value string) (bool, error) {
	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return parsedValue, nil
}

func ConvertInt(value string) (int, error) {
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Could not parse to expected integer value, got: %v", value))
	}
	return parsedValue, nil
}

func ConvertMove(value string) (models.Move, error) {
	// Your conversion logic here.
	valueTrimmed := strings.TrimSpace(value)
	switch strings.ToLower(valueTrimmed) {
	case "rock":
		return models.Rock{}, nil
	case "paper":
		return models.Paper{}, nil
	case "scissor":
		return models.Scissor{}, nil
	default:
		return models.Rock{}, errors.New("Not a valid move: " + valueTrimmed)
	}
}
