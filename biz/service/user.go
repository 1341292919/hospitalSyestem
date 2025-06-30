package service

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/user"
	"Hospital/pkg/errno"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

func (s *UserService) Login(req *user.LoginRequest) (*db.User, error) {
	return db.LoginCheck(s.ctx, req.Username, req.Password, req.Type, req.ID)
}

func (s *UserService) UpdateNurse(req *user.UpdateNurseMessageRequest) error {
	var r *db.UpdateNurseRequest
	id := GetUserIDFromContext(s.c)
	if id > 100 && id != req.ID {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "you have no permission change other doctor")
	}
	if req.Name == nil {
		r = &db.UpdateNurseRequest{
			Position:     req.Position,
			ContactPhone: req.ContactPhone,
			Department:   req.Department,
			Id:           req.ID,
		}
	} else {
		r = &db.UpdateNurseRequest{
			Name:         *req.Name,
			Position:     req.Position,
			ContactPhone: req.ContactPhone,
			Department:   req.Department,
			Id:           req.ID,
		}
	}
	return db.UpdateNurseMessage(s.ctx, r)
}

func (s *UserService) UpdateDoctor(req *user.UpdateDoctorMessageRequest) error {
	var r *db.UpdateDoctorRequest
	id := GetUserIDFromContext(s.c)
	if id > 100 && id != req.ID {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "you have no permission change other doctor")
	}
	if req.Name == nil {
		r = &db.UpdateDoctorRequest{
			Specialty:    req.Specialty,
			Title:        req.Title,
			ContactPhone: req.ContactPhone,
			Department:   req.Department,
			Id:           req.ID,
		}
	} else {
		r = &db.UpdateDoctorRequest{
			Name:         *req.Name,
			Specialty:    req.Specialty,
			Title:        req.Title,
			ContactPhone: req.ContactPhone,
			Department:   req.Department,
			Id:           req.ID,
		}
	}
	return db.UpdateDoctorMessage(s.ctx, r)
}

func (s *UserService) NewUser(req *user.NewUserRequest) (*db.User, error) {
	return db.CreateUser(s.ctx, req.Username, req.Password, req.Type, req.ID)
}

func (s *UserService) QueryUser(req *user.QueryUserRequest) (*db.User, error) {
	return db.GetUserMessage(s.ctx, req.Type, req.ID)
}

func (s *UserService) QueryUserList(req *user.QueryUserListRequest) ([]*db.User, error) {
	switch req.Type {
	case 1: //管理员
		return db.GetAdminList(s.ctx)
	case 2: //医生
		return db.GetDoctorList(s.ctx)
	case 3: //护士
		return db.GetNurseList(s.ctx)
	}
	return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "invalid type")

}
