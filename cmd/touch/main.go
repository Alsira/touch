/*********************************************************************************
*	AUTHOR : Tyler Johnson                                                       *
*	DESCRIPTION : This should work like touch on linux. Made it for some fun.    *
**********************************************************************************/

package main

import (
	"flag"
	"fmt"
	"os"
)

/* This really only makes files right now */
func main() {

	if len(os.Args) < 2 {
		fmt.Errorf("ERROR : Not enough arguments given.\n")
		os.Exit(-1)
	}

	flag.Usage = Usage

	if os.Args[1] == "--help" || os.Args[1] == "/?" {

		flag.Usage()
		os.Exit(0)

	}

	for i := 1; i < len(os.Args); i++ {

		_, err := os.Create(os.Args[i])

		if err != nil {
			fmt.Errorf("ERROR : Cannot create file %s", os.Args[i])
			os.Exit(-1)
		}

	}

	os.Exit(0)

}
