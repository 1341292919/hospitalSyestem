package pack

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/model"
)

func User(data *db.User) *model.User {
	return &model.User{
		ID:           data.Id,
		Name:         data.Name,
		ContactPhone: data.ContactPhone,
		Position:     data.Position,
		Specialty:    data.Specialty,
		Department:   data.Department,
		Title:        data.Title,
		Identity:     data.Identity,
		CreatedAt:    data.CreateTime.Format("2006-01-02 15:04:05"),
	}
}

func UserList(data []*db.User) *model.UserList {
	var users []*model.User
	for _, v := range data {
		users = append(users, User(v))
	}
	return &model.UserList{
		Items: users,
		Total: int64(len(data)),
	}
}
