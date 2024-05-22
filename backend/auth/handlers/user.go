package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/devShahriar/alocmedia/backend/auth/db"
	"github.com/go-playground/validator/v10"
)

type UsersHandler struct {
	l *log.Logger
}

var validate *validator.Validate

func NewUserHandler(l *log.Logger) *UsersHandler {
	return &UsersHandler{l}
}

//Login checker
func (u *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	userinfo := &db.UserLogin{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(userinfo)
	if err != nil {
		http.Error(w, "a", http.StatusBadRequest)
	}
	userResponse, res, err := userinfo.LoginUser()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println(res)
	if res {
		fmt.Println(userResponse)
		expirationTime := time.Now().Add(5 * time.Minute)

		http.SetCookie(w, &http.Cookie{
			Name:     "authT",
			Value:    userResponse,
			Path:     "/",
			Expires:  expirationTime,
			HttpOnly: true,
			Domain:   "http://localhost:3000",
		})
		p, err := json.Marshal(userResponse)
		w.Write([]byte(p))
		if err != nil {
			http.Error(w, "a", http.StatusBadRequest)
		}
	} else {
		Msg := &db.UserResponse{Msg: "invalid user"}
		err = json.NewEncoder(w).Encode(Msg)
		if err != nil {
			http.Error(w, "a", http.StatusBadRequest)
		}

	}

}

//jwt create cookie
//func CreateToken()

//Sign in User Handler
func (u *UsersHandler) InsertUser(w http.ResponseWriter, r *http.Request) {
	user := &db.User{}
	err := user.FromJson(r.Body)

	if err != nil {
		http.Error(w, "unable to parse the json body", http.StatusBadRequest)
	}
	validate = validator.New()
	fmt.Println(user)
	errs := validate.Struct(user)
	if errs != nil {

		u.l.Println(errs)
		http.Error(w, "Data is not correct", http.StatusBadRequest)
	}

	err = user.InsertUser(w)

	if err != nil {
		u.l.Println(err.Error())
		http.Error(w, "User was not inserted ", http.StatusBadRequest)
	}
}
