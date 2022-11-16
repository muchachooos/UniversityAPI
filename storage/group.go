package storage

import (
	"UniversityAPI/model"
)

func (u *UniversityStorage) CreateGroupInDB(id, course, places, spec string) error {

	_, err := u.DataBase.Exec("INSERT INTO `group`(id, course, number_of_places, specialization) VALUES (?,?,?,?)", id, course, places, spec)
	if err != nil {
		return err
	}

	return nil
}

func (u *UniversityStorage) DeleteGroupFromDB(id string) (bool, error) {

	res, err := u.DataBase.Exec("DELETE FROM `group` WHERE id = ?", id)
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

func (u *UniversityStorage) AddToGroupInDB(groupId, studId string) (bool, error) {

	res, err := u.DataBase.Exec("UPDATE student SET group_id = ? WHERE id = ?", groupId, studId)
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

func (u *UniversityStorage) RemoveFromGroupInnDB(groupId, studId string) (bool, error) {

	res, err := u.DataBase.Exec("UPDATE student SET group_id = NULL WHERE group_id = ? AND id = ?", groupId, studId)
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

func (u *UniversityStorage) GetGroupStudentFromDB(groupId string) ([]model.Student, error) {

	var resultTable []model.Student

	err := u.DataBase.Select(&resultTable, "SELECT * FROM student WHERE group_id = ?", groupId)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func (u *UniversityStorage) GetGroupFromDB(groupId string) ([]model.Group, error) {

	var resultTable []model.Group

	err := u.DataBase.Select(&resultTable, "SELECT * FROM `group` WHERE id = ?", groupId)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}
