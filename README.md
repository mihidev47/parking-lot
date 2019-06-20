# PARKING LOT

### Getting started

Ikuti Langkah-langkah Golang + Mysql:

- Copy folder parking-lot ke folder go/src anda.
- Export File parking.sql pada database anda (file disertakan dalam foler).
- Ubah Config DB pada file parking-lot/models/db.go sesuai database anda.

> DB, err = sql.Open("mysql","USERNAME:PASSWORD@/DATABASE_NAME")

Ikuti Perintah Dibawah ini:

Masuk ke Folder

parking_lot/exec/create_parking_lot
> go build create_parking_lot.go
> cp create_parking_lot /usr/local/bin

parking_lot/exec/leave
> go build leave.go
> cp leave /usr/local/bin

parking_lot/exec/park
> go build park.go
> cp park /usr/local/bin

parking_lot/exec/registration_numbers_for_cars_with_colour
> go build registration_numbers_for_cars_with_colour.go
> cp registration_numbers_for_cars_with_colour /usr/local/bin

parking_lot/exec/slot_number_for_registration_number
> go build slot_number_for_registration_number.go
> cp slot_number_for_registration_number /usr/local/bin

parking_lot/exec/slot_numbers_for_cars_with_colour
> go build slot_numbers_for_cars_with_colour.go
> cp slot_numbers_for_cars_with_colour /usr/local/bin

parking_lot/exec/status
> go build status.go
> cp status /usr/local/bin

"Unit Test" dalam folder parking_lot/src/controllers:

> go test -v