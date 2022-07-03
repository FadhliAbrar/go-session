package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/FadhliAbrar/golang-session/model"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(DB *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (database *UserRepositoryImpl) Create(ctx context.Context, user model.User) (model.User, error) {
	query := "INSERT INTO User(Username, Password) VALUES(?,?);"
	db := database.DB
	defer db.Close()
	result, err := db.ExecContext(ctx, query, user.Username, user.Password)
	if err != nil {
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.Id = int(id)
	return user, nil
}

func (database *UserRepositoryImpl) FindUserById(ctx context.Context, Id int) model.User {
	query := "SELECT Username, Password FROM User WHERE Id=?;"
	db := database.DB
	defer db.Close()
	row := db.QueryRowContext(ctx, query, Id)

	user := model.User{}

	err := row.Scan(user.Username, user.Password)
	if err != nil {
		panic(err)
	}

	return user
}

func (database *UserRepositoryImpl) FindUserByUsername(ctx context.Context, username string) (model.User, error) {
	query := "SELECT * FROM User WHERE Username=?;"
	db := database.DB
	defer db.Close()
	rows, err := db.QueryContext(ctx, query, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	user := model.User{}

	if rows.Next() == true {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			panic(err)
		}
	} else {
		return user, errors.New("User tidak ditemukan")
	}

	return user, nil
}
