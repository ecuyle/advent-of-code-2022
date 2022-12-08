package utils

import "fmt"

func CheckError(err error) {
	if err != nil {
		fmt.Println("-------------------------------------------------------")
		fmt.Println("There was an error. See output below.")
		fmt.Println("-------------------------------------------------------")
		panic(err)
	}
}
