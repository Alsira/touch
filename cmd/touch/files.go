package main

/*******************************************************************************************************
*	                                                                                                   *
*	Author : Tyler Johnson                                                                             *
*	Description : Works kinda like a file api. I get kinda lazy so this handles basic operations       *
*	Why : This doesn't really have to exist. It does only cause it makes my life a little easier       *
*                                                                                                      *
********************************************************************************************************/

import (
	"errors"
	"os"
	"time"

	"github.com/djherbis/atime"
)

/**
* * This checks if the filepath given exists
* @param filepath The file path
* @return Returns true if the file exists
*/
func FileExists(filepath string) bool {

	_, err := os.Stat(filepath)
	return errors.Is(err, os.ErrNotExist)

}

/**
* * Gets and returns the access time of the file
* @param filepath The path to the file needing to be checked
* @return Returns the time of last access or a file error if it fails
*/
func GetFileAccessTime(filepath string) (time.Time, error) {

	return atime.Stat(filepath)

}

/**
* * Gets the modification time of the file
* @param filepath The path to the file
* @return Returns the modification time of the file or an error
*/
func GetFileModificationTime(filepath string) (time.Time, error) {

	info, err := os.Stat(filepath)

	if err != nil {
	
		return time.Now(), err
	
	} else {

		return info.ModTime(), nil

	}

}