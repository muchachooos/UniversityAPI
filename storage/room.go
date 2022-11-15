package storage

import (
	"UniversityAPI/model"
	"github.com/jmoiron/sqlx"
)

func CreateRoomInDB(db *sqlx.DB, roomNum, beds string) error {

	_, err := db.Exec("INSERT INTO rooms(room_number, number_of_beds) VALUES (?,?)", roomNum, beds)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRoomFromDB(db *sqlx.DB, roomNum string) (bool, error) {

	res, err := db.Exec("DELETE FROM rooms WHERE room_number = ?", roomNum)
	if err != nil {
		return false, err
	}

	countOfDeletedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if countOfDeletedRows == 0 {
		return false, nil
	}

	return true, nil
}

func AddToRoomInDB(db *sqlx.DB, studId, roomNum string) (bool, error) {

	res, err := db.Exec("UPDATE student SET room = ? WHERE id = ?", studId, roomNum)
	if err != nil {
		return false, err
	}

	countOfModifiedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if countOfModifiedRows == 0 {
		return false, nil
	}

	return true, nil
}

func RemoveFromRoomInDB(db *sqlx.DB, studId, roomNum string) (bool, error) {

	res, err := db.Exec("UPDATE student SET room = NULL WHERE id = ? AND room = ?", studId, roomNum)

	countOfModifiedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if countOfModifiedRows == 0 {
		return false, nil
	}

	return true, nil
}

func GetRoomStudentFromDB(db *sqlx.DB, roomNum string) ([]model.Student, error) {

	var resultTable []model.Student

	err := db.Select(&resultTable, "SELECT * FROM student WHERE room = ?", roomNum)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func GetRoomFromDB(db *sqlx.DB, roomNum string) ([]model.Room, error) {

	var resultTable []model.Room

	err := db.Select(&resultTable, "SELECT * FROM rooms WHERE room_number = ?", roomNum)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}
