package storage

import (
	"UniversityAPI/model"
)

func (u *UniversityStorage) CreateRecordBookInDB(studId string) error {

	_, err := u.DataBase.Exec("INSERT INTO record_book(id_student) VALUES (?)", studId)
	if err != nil {
		return err
	}

	return nil
}

func (u *UniversityStorage) GetIDRecordBookFromDB(studId string) ([]model.ID, error) {

	var resultTable []model.ID

	err := u.DataBase.Select(&resultTable, "SELECT id FROM record_book WHERE id_student = ?", studId)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func (u *UniversityStorage) DeleteRecordBookFromDB(bookId string) (bool, error) {

	res, err := u.DataBase.Exec("DELETE FROM record_book WHERE id = ?", bookId)
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

func (u *UniversityStorage) GetRecordBookFromDB(bookId string) ([]model.RecordBook, error) {

	var resultTable []model.RecordBook

	err := u.DataBase.Select(&resultTable, "SELECT * FROM record_book WHERE id = ?", bookId)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}
