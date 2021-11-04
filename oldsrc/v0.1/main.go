package main

import (
	"fmt"
	"os"
)

func main() {

	if ok, data := handleArg("-v"); ok {
		fmt.Println("Version:", data)
	} else {
		fmt.Println("No version specified")
	}

}

func handleArg(arg string) (bool, string) {
	args := os.Args
	var data string = ""
	for i, v := range args {
		if v == arg {
			if i+1 < len(args) {
				data = args[i+1]
			}
			return true, data
		}
	}
	return false, data
}
