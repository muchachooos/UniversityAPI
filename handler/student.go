package handler

import "github.com/jmoiron/sqlx"

func createStudentInDB(db *sqlx.DB, firstN, lastN, birth string) error {

	_, err := db.Exec("INSERT INTO student(first_name, last_name, date_of_birth) VALUES (?,?,?)", firstN, lastN, birth)
	if err != nil {
		return err
	}

	return nil
}

func getIdStudentFromDB(db *sqlx.DB, firstN, lastN, birth string) ([]ID, error) {

	var resultTable []ID

	err := db.Select(&resultTable, "SELECT id FROM student WHERE first_name = ? AND last_name = ? AND date_of_birth = ?", firstN, lastN, birth)
	if err != nil {
		return nil, err
	}
	return resultTable, nil
}

func deleteStudentFromDB(db *sqlx.DB, id string) (bool, error) {

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

func getStudentFromDB(db *sqlx.DB, id string) ([]Student, error) {

	var resultTable []Student

	err := db.Select(&resultTable, "SELECT * FROM student WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func getStudentByNameFromDB(db *sqlx.DB, firstN, lastN string) ([]Student, error) {

	var resultTable []Student

	err := db.Select(&resultTable, "SELECT * FROM student WHERE first_name = ? AND last_name = ?", firstN, lastN)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}
