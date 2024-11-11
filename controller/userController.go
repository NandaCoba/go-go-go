package controller

import (
	"belajar/db"
	"belajar/schema"
	"belajar/utils"
	"fmt"
)

// pointer
// hal aneh di function belum paham

func GetAll() ([]schema.User, error) {
	var users []schema.User
	checkUsers := db.DB.Find(&users)
	if checkUsers == nil {
		utils.JsonError(nil, 500, "Data response failed")
	}
	return users, nil
}

func GetId(id int) (*schema.User, error) {
	var user *schema.User

	findUser := db.DB.Find(&user, id)

	if findUser == nil {
		fmt.Println("id not found")
	}

	return user, nil
}

func Create(nama string, usia int) (*schema.User, error) {
	user := schema.User{
		Nama: nama,
		Usia: usia,
	}

	if nama == "" {
		utils.JsonError(nil, 500, "nama tidak boleh kosong")
	}
	if &usia == nil {
		utils.JsonError(nil, 500, "usia tidak boleh kosong")
	}

	checkCreate := db.DB.Create(&user)

	if checkCreate == nil {
		utils.JsonError(nil, 500, "failed created new users")
	}

	return &user, nil
}

func Update(id int, nama string, usia int) (*schema.User, error) {
	var user *schema.User

	findUser := db.DB.Find(&user, id)
	if findUser == nil {
		utils.JsonError(nil, 404, "id not found")
	}

	user.Nama = nama
	user.Usia = usia

	checkUpdate := db.DB.Save(&user)

	if checkUpdate == nil {
		utils.JsonError(nil, 500, "update user failed")
	}

	return user, nil
}

func Delete(id int) (*schema.User, error) {
	var user *schema.User

	findUser := db.DB.Find(&user, id)

	if findUser == nil {
		utils.JsonError(nil, 404, "id not found")
	}

	deleteUser := db.DB.Delete(&user, id)

	if deleteUser == nil {
		utils.JsonError(nil, 500, "Failed delete user")
	}

	return user, nil
}
