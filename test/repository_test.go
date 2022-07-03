package test

import (
	"context"
	"fmt"
	"github.com/FadhliAbrar/golang-session/database"
	"github.com/FadhliAbrar/golang-session/model"
	"github.com/FadhliAbrar/golang-session/repository"
	"testing"
)

func TestUserRepository(t *testing.T) {
	db := database.GetDatabase()
	userRepo := repository.NewUserRepositoryImpl(db)

	username := "Fadhli"
	password := "kraken1288"

	user := model.User{
		Username: username,
		Password: password,
	}
	result, err := userRepo.Create(context.Background(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestSessionRepository(t *testing.T) {
	db := database.GetDatabase()
	sessionRepo := repository.NewSessionRepositoryImpl(db)

	mySession := model.Session{
		SessionId:  "asdasdasd",
		IsLoggedIn: true,
		UserId:     1,
	}

	sessionRepo.CreateSession(context.Background(), mySession)

}
