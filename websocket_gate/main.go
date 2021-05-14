package main
//启动方式
//先运行 micro api --address 0.0.0.0:8081  --handler=web --namespace=go.micro.web
//再运行程序

import (
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func hi(w http.ResponseWriter, r *http.Request) {

	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	service := web.NewService(
		web.Name("go.micro.web.websocket"),
	)
	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	// websocket interface
	service.HandleFunc("/websocket", hi)
	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}


}


