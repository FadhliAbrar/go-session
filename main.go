package main

import (
	"context"
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"github.com/FadhliAbrar/golang-session/database"
	"github.com/FadhliAbrar/golang-session/model"
	"github.com/FadhliAbrar/golang-session/repository"
	"github.com/kataras/go-sessions"
	"net/http"
	"text/template"
	"time"
)

type Middleware struct {
	Handler http.Handler
}

type ContextKey string

const ContextUserKey ContextKey = "name"

func (m *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if len(request.Cookies()) > 0 {
		cookies, err := request.Cookie("mysession")
		if err != nil {
			panic(err)
		}
		sessionRepo := repository.NewSessionRepositoryImpl(database.GetDatabase())
		user := sessionRepo.FindSessionUserId(context.Background(), cookies.Value)

		parentRequestContext := context.WithValue(request.Context(), "username", user.Id)
		ctx := context.WithValue(parentRequestContext, "id", user.Username)

		m.Handler.ServeHTTP(writer, request.WithContext(ctx))
	} else {
		m.Handler.ServeHTTP(writer, request)
	}
}

//go:embed views/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "views/*.gohtml"))

func getLogin(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "login.gohtml", "")
}

func postLogin(writer http.ResponseWriter, request *http.Request) {
	db := database.GetDatabase()
	userRepo := repository.NewUserRepositoryImpl(db)

	username := request.PostFormValue("username")
	password := request.PostFormValue("password")

	result, err := userRepo.FindUserByUsername(context.Background(), username)
	if err != nil {
		panic(err)
	}
	if result.Password == password {
		config := sessions.Config{
			Cookie:  "mysession",
			Expires: 5 * time.Minute,
		}
		session := sessions.New(config)
		start := session.Start(writer, request)

		mySession := start.ID()

		sessionRepo := repository.NewSessionRepositoryImpl(database.GetDatabase())
		sess := model.Session{
			SessionId:  mySession,
			UserId:     result.Id,
			IsLoggedIn: true,
		}

		sessionRepo.CreateSession(context.Background(), sess)

		ctx := context.WithValue(request.Context(), "sessionId", sess.SessionId)
		http.Redirect(writer, request.WithContext(ctx), "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
		panic(errors.New("Password salah"))
	}

}

func postLogout(writer http.ResponseWriter, request *http.Request) {

	session := sessions.Start(writer, request)
	key := session.GetString("something")

	session.Clear()
	sessions.Destroy(writer, request)

	fmt.Println(key)

	http.Redirect(writer, request, "/", 302)
}

func getRegister(writer http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(writer, "register.gohtml", "")
}

func postRegister(writer http.ResponseWriter, request *http.Request) {
	db := database.GetDatabase()
	userRepo := repository.NewUserRepositoryImpl(db)

	username := request.PostFormValue("username")
	password := request.PostFormValue("password")

	user := model.User{
		Username: username,
		Password: password,
	}
	id, err := userRepo.Create(context.Background(), user)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	http.Redirect(writer, request, "/login", 302)
}

func home(writer http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("id")
	username := request.Context().Value("username")

	myTemplates.ExecuteTemplate(writer, "main.gohtml", map[string]interface{}{
		"id":       id,
		"username": username,
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/register", getRegister)
	mux.HandleFunc("/post/register", postRegister)
	mux.HandleFunc("/login", getLogin)
	mux.HandleFunc("/post/login", postLogin)
	mux.HandleFunc("/logout", postLogout)

	server := http.Server{
		Handler: &Middleware{
			Handler: mux,
		},
		Addr: "localhost:3000",
	}
	server.ListenAndServe()
}
