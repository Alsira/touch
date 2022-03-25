/*********************************************************************************
*                                                                                *
*	AUTHOR : Tyler Johnson                                                       *
*	DESCRIPTION : This should work like touch on linux. Made it for some fun.    *
*   EDIT : It wasn't fun                                                         *
**********************************************************************************/

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

/* This really only makes files right now */
func main() {

	/* Gotta have some arguments. This just makes eveything easier to leave instead of processing the rest */
	if len(os.Args) < 2 {
		_ = fmt.Errorf("%s", "ERROR : Not enough arguments given.\n")
		os.Exit(-1)
	}

	currentTime := time.Now().Format(time.UnixDate)

	/* These are the flags for the command so far */
	accessTimeFlag := flag.String("a", currentTime, "Change file(s) access time only")
	createFlag := flag.Bool("c", false, "Will not create the file(s) if it already exists.")
	dateFlag := flag.String("d", currentTime, "Parses date given and uses it instead of current time.")
	_ = dateFlag

	flag.Bool("f", false, "IGNORED")

	/* Might need to add the -h, --nodereference flag to the list */

	modifyFlag := flag.String("m", currentTime, "Change only the file(s) modification time.")
	referenceFlag := flag.Bool("r", false, "Changes the time-stamp of file(s) with a reference")
	timeFlag := flag.String("t", currentTime, "Create and set the time of the file.")
	_ = timeFlag

	flag.Parse()

	/* Array of file arguments */
	fileArgs := os.Args[(len(os.Args) - flag.NArg()):]

	if len(fileArgs) < 1 {
		_ = fmt.Errorf("%s", "ERROR : Files not provided.\n")
		os.Exit(-1)
	}

	/* If the create flag is set */
	if *createFlag {

		/* If reference file is used */
		if *referenceFlag {

			/* Check if 2 or more files are there*/

			if len(fileArgs) < 2 {
				_ = fmt.Errorf("%s", "ERROR : Not enough files provided.\n")
				os.Exit(-1)
			}

			referenceFilePath := fileArgs[0]

			refAccessTime, err := GetFileAccessTime(referenceFilePath)
			if err != nil {

				fmt.Errorf("%s", err.Error()+"\n")
				os.Exit(-1)

			}

			refModTime, err := GetFileModificationTime(referenceFilePath)
			if err != nil {

				_ = fmt.Errorf("%s", err.Error()+"\n")
				os.Exit(-1)

			}

			/* If we want to change the access time only */
			if *accessTimeFlag != currentTime {

				/* Loop through the files requiring changes */
				for _, file := range fileArgs[1:] {

					mtime, err := GetFileModificationTime(file)

					if err == nil {

						/* This also returns an error, so if we need error checking here remember this */
						os.Chtimes(file, refAccessTime, mtime)

					}

				} /* End of for loop */

			}

			if *modifyFlag != currentTime {

				/* Loop through files and change their modification times */
				for _, file := range fileArgs[1:] {

					atime, err := GetFileAccessTime(file)

					/* May want to do error checking. IDK */
					if err == nil {

						/* REMEMEBER this also returns an error, so if there is any error checking, this would help */
						os.Chtimes(file, atime, refModTime)

					}

				} /* End of for loop */

				/* If not specified this kinda just changes mod and access time */
			} else {

				for _, file := range fileArgs[1:] {

					atime, err1 := GetFileAccessTime(file)
					mtime, err2 := GetFileModificationTime(file)

					if err1 == nil && err2 == nil {

						/* Again, remeber that this changes both and also returns an error */
						os.Chtimes(file, atime, mtime)

					}

				}

			}

		} else { /* End of reference flag check */

			layout := "Mon Jan 02 2006 15:04:05 GMT-0700"

			// TODO : Kinda got to handle setting the files with reference

			/* Handle access time flag */
			if *accessTimeFlag != currentTime {

				/* This will parse the time */
				time, err := time.Parse(layout, *accessTimeFlag)
				if err != nil {
					fmt.Errorf("%s", err.Error())
				}

				/* Go through each file */
				for _, file := range fileArgs {
					modTime, err := GetFileModificationTime(file)
					if err != nil {
						fmt.Errorf("%s", err.Error())
					}

					err = os.Chtimes(file, time, modTime)
					if err != nil {
						fmt.Errorf("%s", err.Error())
					}
				}

			}

			/* Set modification time */
			if *modifyFlag != currentTime {

				time, err := time.Parse(layout, *modifyFlag)
				if err != nil {
					fmt.Errorf("%s", err.Error())
				}

				for _, file := range fileArgs {

					atime, err := GetFileAccessTime(file)
					if err != nil {
						fmt.Errorf("%s", err.Error())
					}

					err = os.Chtimes(file, atime, time)
					if err != nil {
						fmt.Errorf("%s", err.Error())
					}

				}

			}

		}

	} /* End of create flag check */

	os.Exit(0)

}
