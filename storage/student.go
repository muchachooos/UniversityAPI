package storage

import (
	"UniversityAPI/model"
)

func (u *UniversityStorage) CreateStudentInDB(firstN, lastN, birth string) error {

	_, err := u.DataBase.Exec("INSERT INTO student(first_name, last_name, date_of_birth) VALUES (?,?,?)", firstN, lastN, birth)
	if err != nil {
		return err
	}

	return nil
}

func (u *UniversityStorage) GetIdStudentFromDB(firstN, lastN, birth string) ([]model.ID, error) {

	var resultTable []model.ID

	err := u.DataBase.Select(&resultTable, "SELECT id FROM student WHERE first_name = ? AND last_name = ? AND date_of_birth = ?", firstN, lastN, birth)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func (u *UniversityStorage) DeleteStudentFromDB(id string) (bool, error) {

	res, err := u.DataBase.Exec("DELETE FROM student WHERE id = ?", id)
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

func (u *UniversityStorage) GetStudentFromDB(id string) ([]model.Student, error) {

	var resultTable []model.Student

	err := u.DataBase.Select(&resultTable, "SELECT * FROM student WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func (u *UniversityStorage) GetStudentByNameFromDB(firstN, lastN string) ([]model.Student, error) {

	var resultTable []model.Student

	err := u.DataBase.Select(&resultTable, "SELECT * FROM student WHERE first_name = ? AND last_name = ?", firstN, lastN)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}
