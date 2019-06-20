package services


import (
	"strings"
	"strconv"
	"fmt"
	"../models"
)

func Create_Parking_Services(count_slot int) string {
	db := models.Conn()

	Truncate()

	for i := 1; i <= count_slot; i++ {
		insForm, err := db.Prepare("INSERT INTO parking_slot_detail(slot) VALUES(?)")
		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(i)

	}
	defer db.Close()
	slot_car := strconv.Itoa(count_slot)

	return "Created a parking lot with "+slot_car+" slot \n"
}

func Park_Services(command []string) string {
	db := models.Conn()


	selDB, err := db.Query("SELECT * FROM parking_slot_detail")
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := models.Detail_Slot{}

	// Read all rows from database
	// This time we are going to get only one value, doesn't need the slice
	var slot_car string
	incars := 0

	for selDB.Next() {
		// Store query values on this temporary variables
		var id , slot int
		var status string

		// Scan each row to match the ID and check for errors
		err = selDB.Scan(&id, &slot, &status)
		if err != nil {
			panic(err.Error())
		}

		if status == "kosong" && incars == 0{
			ins, err := db.Prepare("INSERT INTO cars (slot_number, plat_number, color) VALUES(?,?,?)")
			if err != nil {
				panic(err.Error())
			}

			// Execute the prepared SQL, getting the form fields
			ins.Exec(slot, command[1], command[2])
			incars = 1

			upd, err := db.Prepare("UPDATE parking_slot_detail SET status=? WHERE slot=?")
			if err != nil {
				panic(err.Error())
			}

			slot_car = strconv.Itoa(slot)

			// Update row based on hidden form field ID
			upd.Exec("terisi", slot)
		}

		// Get the Scan into the Struct
		n.Id = id
		n.Slot = slot
		n.Status = status

	}
	defer db.Close()

	if incars == 0 {
		return "Sorry, parking lot is full"
	}else{
		return "Allocated slot number: "+slot_car
	}
}

func Leave_Services(command []string) string {
	db := models.Conn()


	selDB, err := db.Query("SELECT * FROM parking_slot_detail")
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := models.Detail_Slot{}

	// Read all rows from database
	// This time we are going to get only one value, doesn't need the slice
	outcars := 0

	for selDB.Next() {
		// Store query values on this temporary variables
		var id , slot int
		var status string

		// Scan each row to match the ID and check for errors
		err = selDB.Scan(&id, &slot, &status)
		if err != nil {
			panic(err.Error())
		}

		if status == "terisi" && outcars == 0{
			delForm, err := db.Prepare("DELETE FROM cars WHERE slot_number=?")
			if err != nil {
				panic(err.Error())
			}

			// Execute the Delete SQL
			delForm.Exec(command[1])
			outcars = 1

			upd, err := db.Prepare("UPDATE parking_slot_detail SET status=? WHERE slot=?")
			if err != nil {
				panic(err.Error())
			}

			// Update row based on hidden form field ID
			upd.Exec("kosong", command[1])
		}

		// Get the Scan into the Struct
		n.Id = id
		n.Slot = slot
		n.Status = status

	}

	defer db.Close()

	if outcars == 0 {
		return "Sorry, parking is Already."
	}else{
		return "Slot number "+command[1]+" is free"
	}
}


func Status_Services() []models.Cars_Slot {
	db := models.Conn()

	fmt.Print("Slot No. \t\t\t\t Registration No \t\t\t\t Colour \n")
	
	selDB, err := db.Query("SELECT * FROM cars")
	if err != nil {
		panic(err.Error())
	}

	n := models.Cars_Slot{}

	res := []models.Cars_Slot{}

	for selDB.Next() {
		// Store query values on this temporary variables
		var id , slot_number int
		var plat_number, color string

		// Scan each row to match the ID and check for errors
		err = selDB.Scan(&id, &slot_number, &plat_number, &color)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		slot_car := strconv.Itoa(slot_number)
		
		fmt.Print(slot_car+" \t\t\t\t\t "+plat_number+" \t\t\t\t\t "+color+" \n")

		n.Id = id
		n.Slot_Number = slot_number
		n.Number_Cars = plat_number
		n.Color = color

		res = append(res, n)
	}

	defer db.Close()

	return res
}

func Registration_numbers_for_cars_with_colour_Services(color string) string {
	db := models.Conn()


	selDB, err := db.Query("SELECT * FROM cars WHERE color=?",color)
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := models.Cars_Slot{}

	res := []string{}
	for selDB.Next() {
		// Store query values on this temporary variables
		var id , slot_number int
		var plat_number, color string

		// Scan each row to match the ID and check for errors
		err = selDB.Scan(&id, &slot_number, &plat_number, &color)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.Id = id
		n.Slot_Number = slot_number
		n.Number_Cars = plat_number
		n.Color = color
		res = append(res, n.Number_Cars)
		
	}

	defer db.Close()

	if len(res) > 0 {
		return strings.Join(res, ", ")
	}else{
		return "Not Found!"
	}

}

func Slot_numbers_for_cars_with_colour_Services(color string) string {
	db := models.Conn()


	selDB, err := db.Query("SELECT * FROM cars WHERE color=?",color)
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := models.Cars_Slot{}

	res := []string{}
	for selDB.Next() {
		// Store query values on this temporary variables
		var id , slot_number int
		var plat_number, color string

		// Scan each row to match the ID and check for errors
		err = selDB.Scan(&id, &slot_number, &plat_number, &color)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.Id = id
		n.Slot_Number = slot_number
		n.Number_Cars = plat_number
		n.Color = color
		slot_car := strconv.Itoa(n.Slot_Number)
		res = append(res, slot_car)
	}

	defer db.Close()

	if len(res) > 0 {
		return strings.Join(res, ", ")
	}else{
		return "Not Found!"
	}

}

func Slot_number_for_registration_number_Services(number_car string) string {
	db := models.Conn()


	selDB, err := db.Query("SELECT * FROM cars WHERE plat_number=?",number_car)
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := models.Cars_Slot{}

	res := []string{}
	for selDB.Next() {
		// Store query values on this temporary variables
		var id , slot_number int
		var plat_number, color string

		// Scan each row to match the ID and check for errors
		err = selDB.Scan(&id, &slot_number, &plat_number, &color)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.Id = id
		n.Slot_Number = slot_number
		n.Number_Cars = plat_number
		n.Color = color
		slot_car := strconv.Itoa(n.Slot_Number)
		res = append(res, slot_car)
	}

	defer db.Close()

	if len(res) > 0 {
		return strings.Join(res, ", ")
	}else{
		return "Not Found!"
	}
	
}

func Truncate()  {
	db := models.Conn()


	_, err := db.Query("TRUNCATE TABLE parking_slot_detail")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("TRUNCATE TABLE cars")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

}