package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro_game/user_service/handler"
	pb "github.com/micro_game/user_service/proto/user"
	repository "github.com/micro_game/user_service/repo"
)

func main() {
	repo := &repository.UserRepository{}

	srv := micro.NewService(
		micro.Name("micrograms.service.user"),
		micro.Version("latest"), // 新增接口版本参数
	)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{Repo: repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
