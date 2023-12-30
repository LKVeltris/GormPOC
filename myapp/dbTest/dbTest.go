package dbTest

import (
	"fmt"
	"myapp/database"
	"myapp/models"
	"myapp/repositories"
)

func DBTest() {

	database.Connect()
	var u1 = models.User{
		ID:             1,
		Name:           "John Doe",
		Email:          "abc@gmail.com",
		OrganizationID: "ABC1",
		Settings:       "setting1",
	}
	var u2 = models.User{
		ID:             2,
		Name:           "Max Doe",
		Email:          "max@gmail.com",
		OrganizationID: "MAX1",
		Settings:       "setting2",
	}

	fmt.Println("=========== CREAT USERS IN DB ======================")
	userRepo := new(repositories.UserRepository)
	a, err := userRepo.CreateUser(&u1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("John : %+v\n", a)

	a, err = userRepo.CreateUser(&u2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Max: %+v\n", a)

	fmt.Println("=========== GET USER BY ID ======================")
	a, err = userRepo.GetUserByID(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", a)

	fmt.Println("=========== Update USER ======================")
	u1.Email = "def@gmail.com"
	err = userRepo.UpdateUser(&u1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("=========== LIST USER ======================")
	ul, err := userRepo.ListUsers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", ul)

	fmt.Println("=========== DELETE USER BY ID ======================")
	err = userRepo.DeleteUser(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("=========== DELETE USER BY ID ======================")
	err = userRepo.DeleteUser(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("=========== LIST USER ======================")
	ul, err = userRepo.ListUsers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", ul)
}
