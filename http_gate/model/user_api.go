package model

import (
	"encoding/json"
	api "github.com/micro/go-micro/api/proto"
	user "github.com/micro_game/user_service/proto/user"
	"golang.org/x/net/context"
	"strconv"
)

type User struct {
	U user.UserServiceClient
}

func (srv *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := srv.U.Login(ctx, &user.Request{})

	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"serverid":    strconv.Itoa(int(response.Serverid))  ,
		"flag":        strconv.Itoa(int(response.Flag)),
		"currenttime": strconv.FormatInt(response.Currenttime,10),
		"accessport":  strconv.Itoa(int(response.AccessPort)),
		"curpt":       strconv.Itoa(int(response.CurPt)),
	})
	rsp.Body = string(b)

	return nil
}
