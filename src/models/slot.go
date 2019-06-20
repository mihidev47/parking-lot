package models

type Slot struct {
	Id int
	Parking_Slot int
}

type Detail_Slot struct {
	Id int
	Slot int
	Status string
}