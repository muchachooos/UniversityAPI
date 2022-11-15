package handler

import "github.com/jmoiron/sqlx"

func createRoomInDB(db *sqlx.DB, roomNum, beds string) error {

	_, err := db.Exec("INSERT INTO rooms(room_number, number_of_beds) VALUES (?,?)", roomNum, beds)
	if err != nil {
		return err
	}

	return nil
}

func deleteRoomFromDB(db *sqlx.DB, roomNum string) (bool, error) {

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

func addToRoomInDB(db *sqlx.DB, studId, roomNum string) (bool, error) {

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

func removeFromRoomInDB(db *sqlx.DB, studId, roomNum string) (bool, error) {

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
