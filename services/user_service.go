package services

import (
	"fmt"

	"github.com/afaf-tech/bookstore_users-api/domain/users"
	"github.com/afaf-tech/bookstore_users-api/utils/errors"
)

/* Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan
tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
*/
/*
 Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*)
 tepat sebelum nama variabel. Metode ini disebut dengan dereferencing. */

// the error return should be in the end
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil

}
func GetUser(userId int64) (*users.User, *errors.RestErr) {

	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	fmt.Println(result) 
	return result, nil
}
func FindUser() {}
