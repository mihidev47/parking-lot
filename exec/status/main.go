package main

import (
	"fmt"
	"os"
	"../../src/controllers"

)

func main()  {
	args := os.Args
	if len(args) != 0 {
		cplValue := os.Args
		cpl := controllers.Command_Parking(cplValue)
		fmt.Println(cpl)
	}
}