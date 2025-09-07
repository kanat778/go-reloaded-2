package main

import (
	"fmt"
	"os"
	proj "reloaded/internal"
)

func main() {
	args := os.Args[1:]
	if err := proj.ValidArgs(args); err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("ERROR: Not found the file!")
		return
	}

	//stringArray := proj.BytesToStrArray(bytes)
	//stringArray = proj.CommandFix(stringArray)
	//stringArray = proj.ArticleFix(stringArray)

	stringArray := proj.FixAll(bytes)
	fmt.Println(stringArray)

	//for ind, str := range stringArray {
	//
	//}
	//res := 0

	t := proj.IsCommand("(backup, 5)")
	fmt.Println(t)

}
