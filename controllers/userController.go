package controllers

import (
	"belajar/db"
	"belajar/schema"
	"fmt"
)

func GetAll() ([]schema.User, error) {
	var user []schema.User

	err := db.DB.Find(&user)

	if err != nil {
		fmt.Println("User not founds")
	}

	return user, nil
}

func Create(name string, age int) (*schema.User, error) {
	user := schema.User{
		Name: name,
		Age:  age,
	}
	err := db.DB.Create(&user)

	if err != nil {
		fmt.Println("Ada error saat create user")
	}

	return &user, nil
}

func Update(id int, name string, age int) (*schema.User, error) {
	user := &schema.User{}
	if checkId := db.DB.Find(user, id).Error; checkId != nil {
		fmt.Println("id not found")
	}
	user.Name = name
	user.Age = age

	updateSave := db.DB.Save(&user)

	if updateSave != nil {
		fmt.Println("save is failed")
	}

	return user, nil
}

func Delete(id int) (*schema.User, error) {
	user := &schema.User{}
	findUser := db.DB.Find(user, id)

	if findUser != nil {
		fmt.Println("user not found")
	}

	deleteUser := db.DB.Delete(user, id)

	if deleteUser != nil {
		fmt.Println("user failed delete")
	}
	return user, nil
}
