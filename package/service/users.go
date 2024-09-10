package service

import (
	"Taxi_service/models"
	"Taxi_service/package/repository"
	"Taxi_service/utils"
	"fmt"
)

func CreateUser(user models.User) error {
	_, err := repository.GetUserByUsernameAndPassword(user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
	}
	user.Password = utils.GenerateHash(user.Password)

	err = repository.CreateUser(user)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func UpdateUser(user models.User, id int) error {
	err := repository.EditUser(user.FullName, user.Username, user.Rating, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func UpdateUserRating(user models.User, id int) error {
	err := repository.EditUserRating(user.Rating, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func IsBlockedUser(isblocked bool, id int) error {
	err := repository.BlockedUser(isblocked, id)
	if err != nil {
		fmt.Println(err)

	}
	return err
}
func IsDeletedUser(isdeleted bool, id int) error {
	err := repository.DeleteUser(isdeleted, id)
	if err != nil {
		fmt.Println(err)

	}
	return err
}

func PrintAllUsers(isdeleted bool, isblocked bool, urole string) (user []models.User, err error) {
	users, err := repository.GetAllUsers(isdeleted, isblocked, urole)
	if err != nil {
		return users, err

	}
	return users, nil
}

func PrintAllUsersByID(isdeleted bool, isblocked bool, id int) (user []models.User, err error) {
	users, err := repository.GetAllUserByID(isdeleted, isblocked, id)
	if err != nil {
		return users, err

	}
	return users, nil
}
