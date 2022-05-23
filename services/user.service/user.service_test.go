package user_service_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	m "mongo-crud-go/models"
	userService "mongo-crud-go/services/user.service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userId string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	userId = oid.Hex()

	random := fmt.Sprint(rand.Intn(100))

	user := m.User{
		ID:        oid,
		Name:      "Oliver" + random,
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

func TestUpdate(t *testing.T) {

	user := m.User{
		Name:  "Smith",
		Email: "smith@mail.com",
	}

	err := userService.Update(user, userId)

	if err != nil {
		t.Error("error update user data - test fail")
		t.Fail()
	}

	t.Log("test completed successfully")
}

func TestDelete(t *testing.T) {
	err := userService.Delete(userId)

	if err != nil {
		t.Error("Error on delete user - test fail")
		t.Fail()
	} else {
		t.Log("test completed successfully")
	}

}
