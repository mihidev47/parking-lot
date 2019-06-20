package controllers

import (
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	total := Create_Slot("5")
	if total != "Created a parking lot with 5 slots" {
		t.Errorf("CreateParkingLot was incorrect, got: %s, want: %s.", total, "Created a parking lot with 5 slots")
	}
}

func TestPark(t *testing.T) {
	//reset
	Create_Slot("5")
	park := Park([]string{"park", "KA-01-HH-1234", "White"})
	if park != "Allocated slot number: 1" {
		t.Errorf("Park was incorrect, got: %s, want: %s.", park, "Allocated slot number: 1")
	}
}

func TestLeave(t *testing.T) {
	leave := Leave([]string{"leave","1"})
	if leave != "Slot number 1 is free" {
		t.Errorf("Leave was incorrect, got: %s, want: %s.", leave, "Slot number 1 is free")
	}
}

func TestStatus(t *testing.T) {
	Create_Slot("5")
	Park([]string{"park", "KA-01-HH-1234", "White"})
	all := Status()
	if len(all) != 1 {
		t.Errorf("Status was incorrect, got: %d, want: %d.", len(all), 1)
	}
}

func TestRegistrationNumbersForCarsWithColour(t *testing.T) {
	carPlateNumber := Registration_numbers_for_cars_with_colour([]string{"registration_numbers_for_cars_with_colour","White"})
	if carPlateNumber != "KA-01-HH-1234" {
		t.Errorf("RegistrationNumbersForCarsWithColour was incorrect, got: %s, want: %s.", carPlateNumber, "KA-01-HH-1234")
	}
}

func TestSlotNumbersForCarsWithColour(t *testing.T) {
	slotNumber := Slot_numbers_for_cars_with_colour([]string{"slot_number_for_registration_number","White"})
	if slotNumber != "1" {
		t.Errorf("SlotNumbersForCarsWithColour was incorrect, got: %s, want: %s.", slotNumber, "1")
	}
}

func TestSlotNumberForRegistrationNumber(t *testing.T) {
	slotNumber := slot_number_for_registration_number([]string{"slot_number_for_registration_number","KA-01-HH-1234"})
	if slotNumber != "1" {
		t.Errorf("Park was incorrect, got: %s, want: %s.", slotNumber, "1")
	}
}