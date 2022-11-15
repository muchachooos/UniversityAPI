package storage

import (
	"UniversityAPI/model"
	"github.com/jmoiron/sqlx"
)

func CreateStudentInDB(db *sqlx.DB, firstN, lastN, birth string) error {

	_, err := db.Exec("INSERT INTO student(first_name, last_name, date_of_birth) VALUES (?,?,?)", firstN, lastN, birth)
	if err != nil {
		return err
	}

	return nil
}

func GetIdStudentFromDB(db *sqlx.DB, firstN, lastN, birth string) ([]model.ID, error) {

	var resultTable []model.ID

	err := db.Select(&resultTable, "SELECT id FROM student WHERE first_name = ? AND last_name = ? AND date_of_birth = ?", firstN, lastN, birth)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func DeleteStudentFromDB(db *sqlx.DB, id string) (bool, error) {

	res, err := db.Exec("DELETE FROM student WHERE id = ?", id)
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

func GetStudentFromDB(db *sqlx.DB, id string) ([]model.Student, error) {

	var resultTable []model.Student

	err := db.Select(&resultTable, "SELECT * FROM student WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func GetStudentByNameFromDB(db *sqlx.DB, firstN, lastN string) ([]model.Student, error) {

	var resultTable []model.Student

	err := db.Select(&resultTable, "SELECT * FROM student WHERE first_name = ? AND last_name = ?", firstN, lastN)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}
