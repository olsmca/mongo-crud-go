package user_service

import (
	u "mongo-crud-go/models"
	userRepository "mongo-crud-go/repositories/user.repository"
)

func Create(user u.User) error {

	err := userRepository.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func Read() (u.Users, error) {

	users, err := userRepository.Read()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func Update(user u.User, userId string) error {

	err := userRepository.Update(user, userId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {

	err := userRepository.Delete(userId)

	if err != nil {
		return err
	}

	return nil
}
