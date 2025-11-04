package main

import (
	"Game/entity"
	"Game/repository/mysql"
	"fmt"
)

func main() {
	testUsereMysqlRepo()
}

func testUsereMysqlRepo() {
	repo := mysql.New()

	createdUser, err := repo.Register(entity.User{
		ID: 0,
		Name: "mamad",
		PhoneNumber: "02222",
	})
	if err != nil {
		fmt.Println("register user", err)
	} else {
		fmt.Println("create user", createdUser)
	}

	// uniqueness check
	isUnique, err := repo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	if err != nil {
		fmt.Println("unique error: ", err)
	}
	fmt.Println("uniqueness: ", isUnique)
}