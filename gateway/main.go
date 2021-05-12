package main

import (
	"github.com/micro/go-micro"
	user "github.com/micro_game/user_service/proto/user"
)

type User struct {
	U user.UserService
}

func main(){
	service := micro.NewService(
		micro.Name("microgames.gateway"),
	)

	service.Init()


}
