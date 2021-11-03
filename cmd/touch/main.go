/*********************************************************************************
*	AUTHOR : Tyler Johnson                                                       *
*	DESCRIPTION : This should work like touch on linux. Made it for some fun.    *
**********************************************************************************/

package main

import (
	"fmt"
	"os"
	"flag"
	"time"
)

/* This really only makes files right now */
func main() {

	/* Gotta have some arguments. This just makes eveything easier to leave instead of processing the rest */
	if len(os.Args) < 2 {
		fmt.Errorf("ERROR : Not enough arguments given.\n")
		os.Exit(-1)
	}

	/* These are the flags for the command so far */
	accessTimeFlag := flag.Duration("a", time.Now(), "Change file(s) access and modification time.")
	modifyFlag := flag.Duration("m", time.Now(), "Change only the file(s) modification time.")
	referenceFlag := flag.Bool("r", false, "Changes the time-stamp of file(s) with a reference")
	timeFlag := flag.Duration("t", time.Now(), "Create and set the time of the file.")
	createFlag := flag.Bool("c", false, "Will not create the file(s) if it already exists.")

	flag.Parse()
	
	/* If the create flag is set */
	if *createFlag {

		// TODO Gotta see if the reference Flag to see if we are using a reference file

	}

	os.Exit(0)

}