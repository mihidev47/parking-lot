package main

import (
	"strings"
    "bufio"
    "fmt"
    "os"
)
type Parking struct {
	Slot int
}

type Cars_Slot struct {
	Number_Cars string
	Color string
}

func main() {
    reader := bufio.NewReader(os.Stdin)
	fmt.Print("== Parking System ==\n")

	for {
		command, _ := reader.ReadString('\n')
		Command_Parking(command)
	}

	bufio.NewReader(os.Stdin).ReadBytes('\n') 
}


func Command_Parking(command string)  {
	command_slice := strings.Split(command, " ")
	if command_slice[0] == "create_parking_lot"{
		var parking Parking
		parking.Slot = 6
		fmt.Print(parking)
		fmt.Printf("%v %d %v ", "Created a parking lot with ", parking.Slot, " slot")

	}else{
		fmt.Print("Error Command Input\n")
	}
}
