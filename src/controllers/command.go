package controllers

import (
	"strconv"
	"../services"
	"../models"
)

var exist_slot int

//Handle Function
func Command_Parking(command []string) string {

	if command[0] == "create_parking_lot"{
		results := Create_Slot(command[1])
		return results
	}else if command[0] == "park"{
		results := Park(command)
		return results
	}else if command[0] == "leave"{
		results := Leave(command)
		return results
	}else if command[0] == "status"{
		Status()
		return "\n"
	}else if command[0] == "registration_numbers_for_cars_with_colour"{
		results := Registration_numbers_for_cars_with_colour(command)
		return results
	}else if command[0] == "slot_numbers_for_cars_with_colour"{
		results := Slot_numbers_for_cars_with_colour(command)
		return results
	}else if command[0] == "slot_number_for_registration_number"{
		results := slot_number_for_registration_number(command)
		return results
	}else{
		results := "Failed Command"+command[0]
		return results
	}
	
}

//create park function
func Create_Slot(count_slot string) string {

	slot_car, _ := strconv.Atoi(count_slot) 
	services.Create_Parking_Services(slot_car)

	return "Created a parking lot with "+count_slot+" slots"
}

//park function
func Park(commend []string) string  {

	location_slot := services.Park_Services(commend)

	return location_slot

}

//leave park function
func Leave(commend []string) string  {
	
	leave := services.Leave_Services(commend)

	return leave
}

//status park function
func Status() []models.Cars_Slot {

	data := services.Status_Services()

	return data
}

//check park function
func Registration_numbers_for_cars_with_colour(command []string) string {
	data := services.Registration_numbers_for_cars_with_colour_Services(command[1])
	return data
}

//check park function
func Slot_numbers_for_cars_with_colour(command []string) string {
	data := services.Slot_numbers_for_cars_with_colour_Services(command[1])
	return data
}

//check park function
func slot_number_for_registration_number(command []string) string  {
	data := services.Slot_number_for_registration_number_Services(command[1])
	return data
}
