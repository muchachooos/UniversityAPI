package storage

import (
	"UniversityAPI/model"
)

func (u *UniversityStorage) CreateRoomInDB(roomNum, beds string) error {

	_, err := u.DataBase.Exec("INSERT INTO rooms(room_number, number_of_beds) VALUES (?,?)", roomNum, beds)
	if err != nil {
		return err
	}

	return nil
}

func (u *UniversityStorage) DeleteRoomFromDB(roomNum string) (bool, error) {

	res, err := u.DataBase.Exec("DELETE FROM rooms WHERE room_number = ?", roomNum)
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

func (u *UniversityStorage) AddToRoomInDB(studId, roomNum string) (bool, error) {

	res, err := u.DataBase.Exec("UPDATE student SET room = ? WHERE id = ?", studId, roomNum)
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

func (u *UniversityStorage) RemoveFromRoomInDB(studId, roomNum string) (bool, error) {

	res, err := u.DataBase.Exec("UPDATE student SET room = NULL WHERE id = ? AND room = ?", studId, roomNum)

	countOfModifiedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if countOfModifiedRows == 0 {
		return false, nil
	}

	return true, nil
}

func (u *UniversityStorage) GetRoomStudentFromDB(roomNum string) ([]model.Student, error) {

	var resultTable []model.Student

	err := u.DataBase.Select(&resultTable, "SELECT * FROM student WHERE room = ?", roomNum)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func (u *UniversityStorage) GetRoomFromDB(roomNum string) ([]model.Room, error) {

	var resultTable []model.Room

	err := u.DataBase.Select(&resultTable, "SELECT * FROM rooms WHERE room_number = ?", roomNum)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}
