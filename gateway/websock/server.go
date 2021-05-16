package websock


import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Server struct {
	upgrader *websocket.Upgrader
}

func (s *Server) Run() error{
	s.upgrader = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return nil
}

func (s *Server) HandleWebsocket(w http.ResponseWriter, r *http.Request) error{
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrade error", err)
		return err
	}




	return nil
}

//构造函数
func NewServer() *Server{
	return &Server{
		upgrader: nil,
	}
}
