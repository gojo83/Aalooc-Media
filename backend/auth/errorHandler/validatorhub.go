package errorHandler

type Validatorhub struct {
	WsConnections map[string]*Connections
	Register      chan Suscribe
	Unregister    chan Suscribe
}

var Hub = Validatorhub{
	WsConnections: make(map[string]*Connections),
	Register:      make(chan Suscribe),
	Unregister:    make(chan Suscribe),
}

func (v *Validatorhub) Run() {

	for {

		select {
		case s := <-v.Register:
			connection := v.WsConnections[s.UserId]
			if connection == nil {
				v.WsConnections[s.UserId] = s.Conn
			}
		case s := <-v.Unregister:
			connection := v.WsConnections[s.UserId]
			if connection != nil {
				delete(v.WsConnections, s.UserId)
			}
		}
	}
}
