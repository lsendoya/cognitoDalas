package application

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/lsendoya/cognitoDalas/internal/user/domain"
	"github.com/lsendoya/cognitoDalas/internal/user/infrastructure/storage"
	"gorm.io/gorm"
	"log"
	"time"
)

func Service(gormDB *gorm.DB, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	user, errCreate := createUserFromEvent(event)
	if errCreate != nil {
		return event, errCreate
	}

	if err := storage.Create(gormDB, user); err != nil {
		return event, err
	}
	log.Println("User registration is successful")

	return event, nil
}

func createUserFromEvent(event events.CognitoEventUserPoolsPostConfirmation) (domain.User, error) {
	var user domain.User
	fmt.Println(event.Request.UserAttributes)
	for key, value := range event.Request.UserAttributes {

		switch key {
		case "email":
			user.UserEmail = value
		case "sub":
			user.UserUUID = value
		case "name":
			user.Name = value
		case "birthdate":
			user.Birthdate = value
		}
	}
	user.IsAdmin = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return user, nil
}
