package internal

import (
	"errors"
	"strings"
)

func ValidArgs(args []string) error {
	if len(args) != 2 {
		return errors.New("ERROR: Invalid number of arguments")
	}
	if args[0] == args[1] {
		return errors.New("ERROR: Invalid filenames")
	}
	if !strings.Contains(args[0], ".txt") || !strings.Contains(args[1], ".txt") {
		return errors.New("ERROR: Invalid datatype of arguments. Please use .txt notation")
	}
	return nil
}
