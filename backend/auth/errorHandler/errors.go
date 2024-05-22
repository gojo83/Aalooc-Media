package errorHandler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/devShahriar/alocmedia/backend/auth/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Suscribe struct {
	UserId string
	Conn   *Connections
}

const writeWait = 10 * time.Second

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(*http.Request) bool { return true },
}

type Payload struct {
	UserId   string `json:"userId"`
	Msgtype  string `json:"msgtype"`
	Data     string `json:"data"`
	ErrorMsg string `json:"errorMsg"`
}

type Validator struct {
	Data chan *Payload
}

var V = Validator{
	Data: make(chan *Payload),
}

type Connections struct {
	ws *websocket.Conn
}

func (c *Connections) Write(v interface{}) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	if err := c.ws.WriteJSON(v); err != nil {
		return err
	}
	return nil
}

func (s Suscribe) WsListerner() {
	c := s.Conn
	defer func() {
		Hub.Unregister <- s
		c.ws.Close()
	}()
	payload := &Payload{
		UserId:   "",
		Msgtype:  "",
		Data:     "",
		ErrorMsg: "",
	}
	for {
		err := c.ws.ReadJSON(&payload)
		fmt.Printf("Got message: %#v\n", payload)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		if payload.Msgtype == "emailValidation" {
			fmt.Println(payload.Msgtype)
			CheckEmail(payload)
		}
		if payload.Msgtype == "phoneValidation" {
			CheckPhone(payload)
		}

	}

}

type Err struct {
	msg string
}

const (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "asd"
	dbname   string = "gopg"
)

func (s *Suscribe) WsWritter() {

	for {
		select {
		case m := <-V.Data:
			fmt.Printf("writter %#v\n", m)
			conn := Hub.WsConnections[m.UserId]

			if err := conn.Write(m); err != nil {
				fmt.Printf("write error %#v", err)
				return
			}
		}
	}

}

type ErrorHandler struct {
	l *log.Logger
}

func NewErrorHandler(l *log.Logger) *ErrorHandler {
	return &ErrorHandler{l}
}
func (e *ErrorHandler) WsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var userId = params["userId"]
	fmt.Printf("userid ::==>> %s", userId)
	ws, err := upgrader.Upgrade(w, r, nil)
	c := &Connections{ws: ws}
	s := Suscribe{userId, c}

	Hub.Register <- s
	if err != nil {
		fmt.Println(err)
		return
	}

	go s.WsListerner()
	go s.WsWritter()
}

func CheckEmail(payload *Payload) {
	db := util.GetConnection(util.Conn{host, port, user, password, dbname})
	email := payload.Data
	fmt.Printf("\ni am the email %s\n", email)
	query := `select email from userinfo where email=$1`

	res, err := db.Query(query, email)
	if err != nil {
		fmt.Println(err)
		return
	}

	//ok := res.Next()
	//fmt.Printf("\nError msg1 %v\n", ok)
	if res.Next() {
		fmt.Printf("\nError msg if block%v\n", res.Next())
		payload.ErrorMsg = "email is used"
		V.Data <- payload
		return
	} else {
		fmt.Printf("\nError msg else block %v\n", res.Next())
		payload.ErrorMsg = "email is not used"
		V.Data <- payload
		return
	}

}

func CheckPhone(payload *Payload) {
	db := util.GetConnection(util.Conn{host, port, user, password, dbname})
	phone := payload.Data
	query := `select phone from userinfo where phone=$1`

	res, err := db.Query(query, phone)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.Next() {
		payload.ErrorMsg = "phone is used"
	} else {
		payload.ErrorMsg = "phone is not used"
	}
	V.Data <- payload
}
