package models

import (
	"myapp/db"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Mahasiswa struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama" validate:"required"`
	Alamat    string `json:"alamat" validate:"required"`
	Telephone string `json:"telephone" validate:"required"`
}

func FetchAllData() (Response, error) {
	var obj Mahasiswa
	var arrobj []Mahasiswa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM contohtable"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telephone)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "Successfully"
	res.Data = arrobj

	return res, nil

}

func StoreMahasiswa(nama string, alamat string, telephone string) (Response, error) {
	var res Response

	v := validator.New()

	mah := Mahasiswa{
		Nama:      nama,
		Alamat:    alamat,
		Telephone: telephone,
	}

	err := v.Struct(mah)

	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT contohtable (nama, alamat, telephone) VALUES (? , ? , ?)"

	stnt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(nama, alamat, telephone)

	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully inserted"
	res.Data = map[string]int64{
		"last_insert_id": lastInsertedId,
	}

	return res, nil
}

func UpdateMahasiswa(id int, nama string, alamat string, telephone string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE contohtable SET nama = ?, alamat = ?, telephone = ? WHERE id = ?"

	stnt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(nama, alamat, telephone, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully Updated"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}

func DeleteMahasiswa(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM contohtable WHERE id = ?"

	stnt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully Deleted"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}
