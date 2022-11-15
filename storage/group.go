package storage

import (
	"UniversityAPI/model"
	"github.com/jmoiron/sqlx"
)

func CreateGroupInDB(db *sqlx.DB, id, course, places, spec string) error {

	_, err := db.Exec("INSERT INTO `group`(id, course, number_of_places, specialization) VALUES (?,?,?,?)", id, course, places, spec)
	if err != nil {
		return err
	}

	return nil
}

func DeleteGroupFromDB(db *sqlx.DB, id string) (bool, error) {

	res, err := db.Exec("DELETE FROM `group` WHERE id = ?", id)
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

func AddToGroupInDB(db *sqlx.DB, groupId, studId string) (bool, error) {

	res, err := db.Exec("UPDATE student SET group_id = ? WHERE id = ?", groupId, studId)
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

func RemoveFromGroupInnDB(db *sqlx.DB, groupId, studId string) (bool, error) {

	res, err := db.Exec("UPDATE student SET group_id = NULL WHERE group_id = ? AND id = ?", groupId, studId)
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

func GetGroupStudentFromDB(db *sqlx.DB, groupId string) ([]model.Student, error) {

	var resultTable []model.Student

	err := db.Select(&resultTable, "SELECT * FROM student WHERE room = ?", groupId)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}
