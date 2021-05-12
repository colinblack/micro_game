package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/micro_game/user_service/model"
	pb "github.com/micro_game/user_service/proto/user"
	"github.com/micro_game/user_service/repo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)
type UserService struct {
	Repo repo.Repository
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	userItems := make([]*pb.User, len(users))
	for index, user := range users {
		userItem, _ := user.ToProtobuf()
		userItems[index] = userItem
	}
	res.Users = userItems
	return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	// 对密码进行哈希加密
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	userModel := &model.User{}
	user, _ := userModel.ToORM(req)
	if err := srv.Repo.Create(user); err != nil {
		return err
	}
	res.User, _ = user.ToProtobuf()
	return nil
}
