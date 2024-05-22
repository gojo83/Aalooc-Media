package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/devShahriar/alocmedia/backend/auth/util"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
)

type User struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"gte=8,required"`
	Phone       int    `json:"phone" validate:"required"`
	CompanyName string `json:"company" validate:"required"`
}

type ValidInfo struct {
	EmailIsUsed chan bool
	PhoneIsUsed chan bool
}

var ValidUserData = ValidInfo{
	EmailIsUsed: make(chan bool),
	PhoneIsUsed: make(chan bool),
}

const (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "asd"
	dbname   string = "gopg"
)

func (u *User) FromJson(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(u)
}

func (u *User) ToJson(w http.ResponseWriter) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(u)
}

//Inserts user data in the database for sign up
func (u *User) InsertUser(w http.ResponseWriter) error {
	db := util.GetConnection(util.Conn{host, port, user, password, dbname})

	query := `insert into userinfo (name,email,password) values ($1,$2,$3)`

	res, err := db.Exec(query, u.Name, u.Email, u.Password)
	fmt.Println(res)
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}

func (u *UserLogin) AuthorizeUser() {

}

//Login

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:password validate:"required"`
}

type UserResponse struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Msg    string `json:"msg"`
}

func (u *UserLogin) LoginUser() (string, bool, error) {
	db := util.GetConnection(util.Conn{host, port, user, password, dbname})
	query := `select user_id,name,email from userinfo where email=$1 and password=$2`
	res, err := db.Query(query, u.Email, u.Password)
	userResponse := UserResponse{}

	userExist := res.Next()
	var jwt_token string
	if userExist {
		err := res.Scan(&userResponse.UserId, &userResponse.Name, &userResponse.Email)
		if err != nil {

		}
		jwt_token, err = CreateToken(&userResponse)

	}
	if err != nil {
		return "", false, err
	}
	return jwt_token, userExist, nil
}

func CreateToken(u *UserResponse) (string, error) {

	os.Setenv("ACCESS_SECRET", "asdfasdfasdf")
	atClaims := jwt.MapClaims{}
	atClaims["userId"] = u.UserId
	atClaims["name"] = u.Name
	atClaims["email"] = u.Email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
