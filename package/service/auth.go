package service

import (
	"Taxi_service/package/repository"
	"Taxi_service/utils"
)

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}

	accessToken, err = GenerateToken(uint(user.ID), user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
