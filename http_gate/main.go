package main

import (
	"fmt"
	"github.com/micro/go-micro"
	m "github.com/micro_game/http_gate/model"
	user "github.com/micro_game/user_service/proto/user"
)

/*type User struct {
	U user.UserServiceClient
}

func (srv *User) GetAll(ctx context.Context, req *user.Request, rsp *user.Response) error{
	rsp, err := srv.U.GetAll(ctx, req)
	if err != nil{
		return err
	}
	return nil
}

func (srv *User) Create(ctx context.Context, req *user.User, rsp *user.Response) error{
	rsp, err := srv.U.Create(ctx, req)
	if err != nil{
		return err
	}
	return nil
}*/

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.http-gate"), //这里服务名前缀必须是go.micro.api
	)

	service.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&m.User{U: user.NewUserServiceClient("micrograms.service.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		fmt.Println("http_gate error")
	}
}
