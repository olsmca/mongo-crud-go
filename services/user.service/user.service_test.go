package user_service_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	m "mongo-crud-go/models"
	userService "mongo-crud-go/services/user.service"
)

func TestCreate(t *testing.T) {

	user := m.User{
		Name:      "Oliver",
		Email:     "user@mail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := userService.Create(user)

	if err != nil {
		fmt.Println(err)
		t.Error("user data persistence test fails")
		t.Fail()
	} else {
		t.Log("test completed successfully")
	}
}

func TestRead(t *testing.T) {
	users, err := userService.Read()

	if err != nil {
		t.Error("error reading user data - test fail")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("the query does not return users")
	} else {

		for _, us := range users {
			str, err := json.Marshal(us)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
			fmt.Println("[")
			fmt.Println(string(str))
			fmt.Println("]")
		}

		t.Log("test completed successfully")
	}
}

func Update(t *testing.T) {

	user := m.User{
		Name:  "Smith",
		Email: "smith@mail.com",
	}

	err := userService.Update(user, "6287baca7b17aed8a6253c52")

	if err != nil {
		t.Error("error update user data - test fail")
		t.Fail()
	}

	t.Log("test completed successfully")
}
