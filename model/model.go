package model

type Student struct {
	ID          int     `db:"id" json:"id"`
	FirstName   string  `db:"first_name" json:"firstName"`
	LastName    string  `db:"last_name" json:"lastName"`
	GroupID     *string `db:"group_id" json:"GroupID"`
	Room        *string `db:"room" json:"room"`
	DateOfBirth *string `db:"date_of_birth" json:"date_of_birth"`
}

type Room struct {
	RoomNumber string `db:"room_number" json:"room_umber"`
	NumOfDBeds int    `db:"number_of_beds" json:"number_of_beds"`
}

type ID struct {
	ID int `db:"id" json:"id"`
}

type Group struct {
	ID             string `db:"id" json:"id"`
	Course         int    `db:"course" json:"course"`
	NumberOfPlaces int    `db:"number_of_places" json:"number_of_places"`
	Specialization string `db:"specialization" json:"specialization"`
}
