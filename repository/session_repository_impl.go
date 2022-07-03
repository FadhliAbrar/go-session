package repository

import (
	"context"
	"database/sql"
	"github.com/FadhliAbrar/golang-session/model"
)

type SessionRepositoryImpl struct {
	Db *sql.DB
}

func NewSessionRepositoryImpl(db *sql.DB) SessionRepository {
	return &SessionRepositoryImpl{Db: db}
}

func (impl *SessionRepositoryImpl) CreateSession(ctx context.Context, session model.Session) {
	query := "INSERT INTO Session(session_id, IsLoggedIn, UserId) VALUES (?,?,?)"
	db := impl.Db
	defer db.Close()

	_, err := db.ExecContext(ctx, query, session.SessionId, session.IsLoggedIn, session.UserId)
	if err != nil {
		panic(err)
	}
}

func (impl *SessionRepositoryImpl) FindSessionUserId(ctx context.Context, session string) model.User {
	query := "SELECT User.Id, User.Username FROM Session INNER JOIN User ON Session.UserId=User.Id WHERE Session.session_id = ?;"
	db := impl.Db
	defer db.Close()

	row := db.QueryRowContext(ctx, query, session)

	user := model.User{}

	err := row.Scan(&user.Id, &user.Username)
	if err != nil {
		panic(err)
	}
	return user
}
