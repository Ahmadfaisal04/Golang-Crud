package models

import (
	"fmt"
	"pasien/config"
	"pasien/entities"
)

//Read Data

func GetAll() []entities.User {
	query, err := config.DB.Query("SELECT * FROM pasien")

	if err != nil {
		panic(err)
	}

	defer query.Close()

	var pasien []entities.User

	for query.Next() {
		var pasiens entities.User

		if err := query.Scan(&pasiens.Id, &pasiens.NoIdentitas, &pasiens.Nama, &pasiens.JenisKelamin, &pasiens.Alamat, &pasiens.NoHP); err != nil {
			panic(err)
		}

		pasien = append(pasien, pasiens)
	}

	return pasien
}

//Create Data

func Create(pasiens entities.User) bool {
	res, err := config.DB.Exec("INSERT INTO pasien (NoIdentitas, Nama, JenisKelamin, Alamat, NoHP) VALUES (?, ?, ?, ?, ?)", pasiens.NoIdentitas, pasiens.Nama, pasiens.JenisKelamin, pasiens.Alamat, pasiens.NoHP)

	if err != nil {
		panic(err)
	}

	lastinsertid, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastinsertid > 0

}

// Update Data
func Update(id int, pasiens entities.User) bool {
	fmt.Print("ex queru")
	res, err := config.DB.Exec("UPDATE pasien SET NoIdentitas = ?, Nama = ?, JenisKelamin = ?, Alamat = ?, NoHP = ? WHERE Id = ?", pasiens.NoIdentitas, pasiens.Nama, pasiens.JenisKelamin, pasiens.Alamat, pasiens.NoHP, id)

	if err != nil {
		fmt.Print("panic err")
		panic(err)
	}

	rowsaffected, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rowsaffected > 0
}

// Details Data
func Details(id int) entities.User {
	query, err := config.DB.Query("SELECT * FROM pasien WHERE Id = ?", id)

	if err != nil {
		panic(err)
	}

	defer query.Close()

	var pasiens entities.User

	for query.Next() {
		if err := query.Scan(&pasiens.Id, &pasiens.NoIdentitas, &pasiens.Nama, &pasiens.JenisKelamin, &pasiens.Alamat, &pasiens.NoHP); err != nil {
			panic(err)
		}
	}

	return pasiens

}

// Delete Data
func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM pasien WHERE Id = ?", id)

	return err
}
