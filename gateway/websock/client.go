package websock

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/util/log"
)

type Client struct {
	id string
	server *Server
	conn *websocket.Conn
	wBuffer chan []byte //写缓冲
}


func NewClient(server *Server, conn *websocket.Conn) *Client{
	c := &Client{
		id : uuid.New().String(),  //分布式全局唯一id
		server: server,
		conn: conn,
		wBuffer: make(chan []byte, 100),  //如果channel中缓冲数据>100则会阻塞
	}

	go func() {
		//读取写缓冲中的内容

		for b := range c.wBuffer{
			if b == nil {
				continue
			}

			err := conn.WriteMessage(websocket.BinaryMessage, b)
			if err != nil {
				log.Warn("Write client message erro", err)
				break
			}
		}
	}()



}





