package main


import (
	"os"
	"errors"
)

func FileExists(filepath string) bool {

	_, err := os.Stat(filepath)
	return errors.Is(err, os.ErrNotExist)

}