package model

//login -> store session ke db session
//bikin db session, di dalamnya mereference pada user
// di middleware, dicheck. apakah session ada user dengan id blablabla, dan melakukan loggedin.
//kalau tidak jangan lanjut
//kalau iya,

type User struct {
	Id       int
	Username string
	Password string
}
